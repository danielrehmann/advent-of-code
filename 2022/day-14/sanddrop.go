package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WorldItem int

const (
	Air WorldItem = iota
	Sand
	Rock
)

func ParseInputFromFile(f string) (World, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	scan := bufio.NewScanner(file)
	rocks := make([]Coordinate, 0)
	for scan.Scan() {
		rocks = append(rocks, ParseLine(scan.Text())...)
	}
	return toRockMap(rocks), nil
}

func toRockMap(rocks []Coordinate) World {
	allRocks := make(World)
	for _, c := range rocks {
		allRocks[c] = Rock
	}
	return allRocks
}

type World map[Coordinate]WorldItem

func (w World) String() string {
	e := ExtremesOfWorld(w)
	var strb strings.Builder
	for y := e.ymin; y <= e.ymax; y++ {
		for x := e.xmin; x <= e.xmax; x++ {
			switch w[Coordinate{x, y}] {
			case Air:
				strb.WriteRune('.')
			case Rock:
				strb.WriteRune('#')
			case Sand:
				strb.WriteRune('~')
			}
		}
		strb.WriteString("\n")
	}

	return strb.String()
}

func DropNextSand(world World, dropPos Coordinate, extremes ext) bool {
	if world[dropPos] != Air {
		return true
	}
	currentPos, status := dropPos, couldDrop
	for couldDrop == status {
		currentPos, status = hasFreeSpot(world, currentPos, dropPos, extremes.ymax)
	}
	world[currentPos] = Sand
	return status == noWayToGo
}

func DropNextSandWithFloor(world World, dropPos Coordinate, floorY int) bool {
	if world[dropPos] != Air {
		return true
	}
	currentPos, status := dropPos, couldDrop
	for couldDrop == status {
		currentPos, status = hasFreeSpotWithFloor(world, currentPos, dropPos, floorY)
		if status == droppedOnFloor {
			world[Coordinate{currentPos.x, currentPos.y + 1}] = Rock
			status = noWayToGo
		}
		if currentPos == dropPos {
			status = reachedTop
		}
	}
	world[currentPos] = Sand
	return status == noWayToGo
}

type status int

const (
	couldDrop status = iota
	noWayToGo
	endlessPit
	droppedOnFloor
	reachedTop
)

func hasFreeSpot(w World, pos, startingPos Coordinate, lowestRock int) (Coordinate, status) {
	if lowestRock <= pos.y {
		return Coordinate{pos.x, pos.y + 1}, endlessPit
	}
	if down := w[Coordinate{pos.x, pos.y + 1}]; down == Air {
		return Coordinate{pos.x, pos.y + 1}, couldDrop
	}

	if downleft := w[Coordinate{pos.x - 1, pos.y + 1}]; downleft == Air {
		return Coordinate{pos.x - 1, pos.y + 1}, couldDrop
	}

	if downright := w[Coordinate{pos.x + 1, pos.y + 1}]; downright == Air {
		return Coordinate{pos.x + 1, pos.y + 1}, couldDrop
	}
	return pos, noWayToGo
}
func hasFreeSpotWithFloor(w World, pos, startingPos Coordinate, floorLevel int) (Coordinate, status) {
	if floorLevel == pos.y+1 {
		return Coordinate{pos.x, pos.y}, droppedOnFloor
	}
	if down := w[Coordinate{pos.x, pos.y + 1}]; down == Air {
		return Coordinate{pos.x, pos.y + 1}, couldDrop
	}

	if downleft := w[Coordinate{pos.x - 1, pos.y + 1}]; downleft == Air {
		return Coordinate{pos.x - 1, pos.y + 1}, couldDrop
	}

	if downright := w[Coordinate{pos.x + 1, pos.y + 1}]; downright == Air {
		return Coordinate{pos.x + 1, pos.y + 1}, couldDrop
	}
	return pos, noWayToGo
}

func DropAllPossibleSand(w World, startingPos Coordinate) int {
	lr := ExtremesOfWorld(w)
	count := 0
	con := true
	for con {
		count++
		con = DropNextSand(w, startingPos, lr)
		if count%500 == 0 {
			fmt.Println(count)
			fmt.Println(w.String())
		}
		if !con {
			count--
		}
	}

	fmt.Println(count)
	fmt.Println(w.String())
	return count
}

func DropAllPossibleSandWithFloor(w World, startingPos Coordinate) int {
	lr := ExtremesOfWorld(w)
	count := 0
	con := true
	for con {
		count++
		con = DropNextSandWithFloor(w, startingPos, lr.ymax+2)
		if count%200 == 0 {
			fmt.Println(count)
			fmt.Println(w.String())
		}
	}
	fmt.Println(count)
	fmt.Println(w.String())
	return count
}

type ext struct {
	xmin, xmax, ymin, ymax int
}

func ExtremesOfWorld(allRocks World) ext {
	xmax, xmin, ymax, ymin := 500, 500, 0, 0
	for v := range allRocks {
		if v.x < xmin {
			xmin = v.x
		}
		if v.y < ymin {
			ymin = v.y
		}
		if v.x > xmax {
			xmax = v.x
		}
		if v.y > ymax {
			ymax = v.y
		}
	}
	return ext{xmin, xmax, ymin, ymax}
}

func ParseLine(s string) []Coordinate {
	rocksInLine := make([]Coordinate, 0)
	if s == "" {
		return rocksInLine
	}
	edges := strings.Split(s, " -> ")
	var previousEdge *Coordinate
	for _, v := range edges {
		rock := ParseCoordinate(v)
		if previousEdge != nil {
			ribt := rocksInBetween(*previousEdge, rock)
			rocksInLine = append(rocksInLine, ribt...)
		}
		rocksInLine = append(rocksInLine, rock)
		previousEdge = &rock
	}
	return rocksInLine
}

func rocksInBetween(previousEdge Coordinate, rock Coordinate) []Coordinate {
	xdiff, ydiff := previousEdge.x-rock.x, previousEdge.y-rock.y
	if xdiff != 0 && ydiff != 0 {
		panic("Diagonal edges not supported")
	}
	rocks := make([]Coordinate, 0)
	var xch int
	switch {
	case xdiff > 0:
		xch = -1
	case xdiff < 0:
		xch = 1
	}
	for i := xdiff + xch; i != 0; i = i + xch {
		rocks = append(rocks, Coordinate{previousEdge.x - i, previousEdge.y})
	}
	var ych int
	switch {
	case ydiff > 0:
		ych = -1
	case ydiff < 0:
		ych = 1
	}
	for i := ydiff + ych; i != 0; i = i + ych {
		rocks = append(rocks, Coordinate{previousEdge.x, previousEdge.y - i})
	}
	return rocks
}

func ParseCoordinate(v string) Coordinate {
	xy := strings.Split(v, ",")
	if len(xy) != 2 {
		panic("Not a coordinate")
	}
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		panic("x could not be parsed")
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		panic("y could not be parsed")
	}
	return Coordinate{x, y}
}

type Coordinate struct {
	x, y int
}
