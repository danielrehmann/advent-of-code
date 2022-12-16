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

type Grid struct {
	scans []Scan
}

type Scan struct {
	sensor, nextBeacon Pos
	dist               int
}
type gridLimits struct {
	xmin, xmax, ymin, ymax int
}

func FindFreePos(g *Grid, upperLimit int) Pos {
	for y := 0; y <= upperLimit; y++ {
		for x := 0; x <= upperLimit; {
			if skpbl := g.findSkippableX(Pos{x, y}); skpbl > 0 {
				x = x + skpbl
			} else {
				return Pos{x, y}
			}
		}
	}
	panic("Found no unscanned positions in scan area")

}

func (g *Grid) findSkippableX(p Pos) int {
	grid := *g
	for _, s := range grid.scans {
		if d := findDistance(s.sensor, p); d <= s.dist {
			return s.sensor.x + s.dist - abs(p.y-s.sensor.y) - p.x + 1
		}
	}
	return 0
}

func (g *Grid) findGridLimits() gridLimits {
	xmin, xmax, ymin, ymax := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for _, p := range g.scans {
		if p.sensor.x-p.dist < xmin {
			xmin = p.sensor.x - p.dist
		}
		if p.sensor.x+p.dist > xmax {
			xmax = p.sensor.x + p.dist
		}
		if p.sensor.y-p.dist < ymin {
			ymin = p.sensor.y - p.dist
		}
		if p.sensor.y+p.dist > ymax {
			ymax = p.sensor.y + p.dist
		}
	}
	return gridLimits{xmin, xmax, ymin, ymax}
}

func (g *Grid) String() string {
	area := g.findGridLimits()
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("xmin: %d, xmax: %d, ymin: %d, ymax: %d\n", area.xmin, area.xmax, area.ymin, area.ymax))
	for y := area.ymin; y <= area.ymax; y++ {
		for x := area.xmin; x <= area.xmax; x++ {
			switch g.GetItem(Pos{x, y}) {
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

func (g *Grid) GetItem(p Pos) Item {
	grid := *g
	current := Unchecked
	for _, s := range grid.scans {
		switch {
		case p == s.sensor:
			current = Sensor
		case p == s.nextBeacon:
			current = Beacon
		case In(p, s.sensor, s.dist):
			current = Scanned
		}
	}
	return current
}

func In(sensor, p Pos, dist int) bool {
	return findDistance(sensor, p) <= dist
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

func LoadFromFile(fn string) Grid {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	grid := make([]Scan, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, getScanFromLine(scanner.Text()))
	}
	return Grid{grid}
}

func BlockedPositionsInRow(g *Grid, y int) int {
	area := g.findGridLimits()
	rowContent := make([]Item, 0)
	for x := area.xmin; x <= area.xmax; x++ {
		if item := g.GetItem(Pos{x, y}); item == Scanned {
			rowContent = append(rowContent, item)
		}
	}
	return len(rowContent)
}

func NewScan(sensor, nextBeacon Pos) Scan {
	dist := findDistance(sensor, nextBeacon)
	return Scan{sensor, nextBeacon, dist}
}

func getScanFromLine(line string) Scan {
	regexGroups := lineRegex.FindStringSubmatch(line)
	s, nb := Pos{asInt(regexGroups[1]), asInt(regexGroups[2])}, Pos{asInt(regexGroups[3]), asInt(regexGroups[4])}
	return NewScan(s, nb)
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
