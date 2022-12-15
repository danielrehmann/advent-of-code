package day15

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Grid map[Pos]Item

func (g Grid) String() string {
	xmin, xmax, ymin, ymax := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for p := range g {
		if p.x < xmin {
			xmin = p.x
		}
		if p.x > xmax {
			xmax = p.x
		}
		if p.y < ymin {
			ymin = p.y
		}
		if p.y > ymax {
			ymax = p.y
		}
	}
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("xmin: %d, xmax: %d, ymin: %d, ymax: %d\n", xmin, xmax, ymin, ymax))
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			switch g[Pos{x, y}] {
			case Unchecked:
				builder.WriteRune('.')
			case Sensor:
				builder.WriteRune('S')
			case Beacon:
				builder.WriteRune('B')
			case Scanned:
				builder.WriteRune('#')
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type Pos struct {
	x, y int
}

type Item int

const (
	Unchecked Item = iota
	Sensor
	Beacon
	Scanned
)

var lineRegex, _ = regexp.Compile(`Sensor at x=(.?\d+), y=(.?\d+): closest beacon is at x=(.?\d+), y=(.?\d+)`)

func LoadFromFile(fn string, row int) Grid {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	grid := make(Grid, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		addLineToGrid(&grid, scanner.Text(), row)
	}
	return grid
}

func BlockedPositionsInRow(g *Grid, y int) int {
	grid := *g
	rowContent := make([]Item, 0)
	for p := range grid {
		if p.y == y {
			i := grid[p]
			if i == Scanned {
				rowContent = append(rowContent, i)
			}
		}
	}
	return len(rowContent)
}

func addLineToGrid(gridPointer *Grid, line string, row int) *Grid {
	regexGroups := lineRegex.FindStringSubmatch(line)
	s, nb := Pos{asInt(regexGroups[1]), asInt(regexGroups[2])}, Pos{asInt(regexGroups[3]), asInt(regexGroups[4])}
	AddToGrid(gridPointer, s, nb, row)
	return gridPointer
}

func AddToGrid(gridPointer *Grid, s, nb Pos, row int) {
	grid := *gridPointer
	dtb := findDistance(s, nb)
	grid[s] = Sensor
	grid[nb] = Beacon

	for x := -dtb; x <= dtb; x++ {
		if abs(x)+abs(s.y-row) <= dtb {
			currentPos := Pos{x + s.x, row}
			if grid[currentPos] == Unchecked {
				grid[currentPos] = Scanned
			}
		}
	}

}

func findDistance(s, nb Pos) int {
	return abs(s.x-nb.x) + abs(s.y-nb.y)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func asInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
