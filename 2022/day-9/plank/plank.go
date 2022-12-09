package plank

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	R direction = iota
	L           = iota
	U           = iota
	D           = iota
)

func Test() string {
	return "module"
}

type Move struct {
	dir   direction
	steps int
}

type field struct {
	Pos        []pos
	visitedByT map[pos]bool
}

func (f field) VisitedByTCount() int {
	return len(f.visitedByT)
}

type pos struct {
	x, y int
}

func (move Move) ApplyTo(f field) {
	for i := 0; i < move.steps; i++ {
		f.apply(move.dir)
	}
}

func (f field) apply(move direction) {
	switch move {
	case L:
		f.Pos[0] = pos{f.Pos[0].x - 1, f.Pos[0].y}
	case R:
		f.Pos[0] = pos{f.Pos[0].x + 1, f.Pos[0].y}
	case U:
		f.Pos[0] = pos{f.Pos[0].x, f.Pos[0].y + 1}
	case D:
		f.Pos[0] = pos{f.Pos[0].x, f.Pos[0].y - 1}
	}

	for i := 1; i < len(f.Pos); i++ {
		previous := f.Pos[i-1]
		current := f.Pos[i]

		if !current.neighbors(previous) {
			switch dx, dy := previous.x-current.x, previous.y-current.y; {
			case dx != 0 && dy != 0:
				var cx, cy int
				if dx > 0 {
					cx = 1
				} else {
					cx = -1
				}
				if dy > 0 {
					cy = 1
				} else {
					cy = -1
				}
				f.Pos[i] = pos{current.x + cx, current.y + cy}
			case dx > 1:
				f.Pos[i] = pos{current.x + 1, current.y}
			case dx < -1:
				f.Pos[i] = pos{current.x - 1, current.y}
			case dy > 1:
				f.Pos[i] = pos{current.x, current.y + 1}
			case dy < -1:
				f.Pos[i] = pos{current.x, current.y - 1}
			default:
				panic("No neighbor with no change")
			}
		}
	}
	f.visitedByT[f.Pos[len(f.Pos)-1]] = true
}

func (x pos) neighbors(y pos) bool {
	area := make(map[pos]bool, 9)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			area[pos{y.x + i, y.y + j}] = true
		}
	}
	_, found := area[x]
	return found
}

func InitialField() field {
	positions := []pos{{0, 0}, {0, 0}}
	f := field{Pos: positions, visitedByT: map[pos]bool{}}
	f.visitedByT[f.Pos[len(f.Pos)-1]] = true
	return f
}

func Initial10KnotField() field {
	f := field{Pos: []pos{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}, visitedByT: map[pos]bool{}}
	f.visitedByT[f.Pos[len(f.Pos)-1]] = true
	return f
}

func Parse(lines []string) ([]Move, error) {
	moves := make([]Move, 0, len(lines))
	for _, l := range lines {
		parts := strings.Split(l, " ")
		var dir direction
		switch parts[0] {
		case "R":
			dir = R
		case "L":
			dir = L
		case "U":
			dir = U
		case "D":
			dir = D
		default:
			return nil, fmt.Errorf("unknown direction %s", parts[0])
		}

		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		moves = append(moves, Move{dir: dir, steps: steps})
	}
	return moves, nil
}

func InputLineByLine(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic("Could not open file " + filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 100)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
