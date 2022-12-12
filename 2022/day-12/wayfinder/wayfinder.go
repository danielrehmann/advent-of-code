package day12

import (
	"math"
	"sort"

	"golang.org/x/exp/slices"
)

func ParseMap(lines []string) (Pos, Pos, map[Pos]int, Dimensions) {
	var start, end Pos
	fullMap := make(map[Pos]int)
	linecount := len(lines)
	mapDimensions := Dimensions{len(lines[0]), linecount}
	for y, v := range lines {
		for x, r := range v {
			value, t := valueOf(r)
			pos := Pos{x, linecount - y - 1}
			fullMap[pos] = value
			switch t {
			case Start:
				start = pos
			case End:
				end = pos
			}
		}
	}

	return start, end, fullMap, mapDimensions
}

func FindAllVertices(world map[Pos]int, dim Dimensions) []Vertex {
	allVertexes := make(map[Vertex]bool, 0)
	for x := 0; x < dim.x; x++ {
		for y := 0; y < dim.y; y++ {
			from := Pos{x, y}
			for _, vertex := range findPossibleMoves(from, world) {
				allVertexes[vertex] = true
			}
		}
	}
	keys := make([]Vertex, 0, len(allVertexes))
	for v := range allVertexes {
		keys = append(keys, v)
	}

	return keys
}

func FindPath(from, to Pos, vertices []Vertex) []Pos {
	fastAccess := make(map[Pos][]Pos, 0)
	for _, v := range vertices {
		fa, e := fastAccess[v.From]
		if !e {
			fastAccess[v.From] = []Pos{v.to}
		} else {
			fastAccess[v.From] = append(fa, v.to)
		}
	}

	dist := make(map[Pos]int, 0)
	prev := make(map[Pos]*Pos, 0)
	queueMap := make(map[Pos]bool, len(vertices))
	for _, v := range vertices {
		if from == v.From {
			dist[v.From] = 0
		} else {
			dist[v.From] = math.MaxInt
		}
		prev[v.From] = nil
		queueMap[v.From] = true
		if from == v.to {
			dist[v.to] = 0
		} else {
			dist[v.to] = math.MaxInt
		}
		prev[v.to] = nil
		queueMap[v.to] = true
	}
	queue := make([]Pos, 0, len(queueMap))
	for p := range queueMap {
		queue = append(queue, p)
	}

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return dist[queue[i]] < dist[queue[j]]
		})
		minPos := queue[0]
		queue = queue[1:]
		for _, v := range fastAccess[minPos] {
			if slices.Contains(queue, v) {
				temp := dist[minPos] + 1
				if temp < dist[v] {
					dist[v] = temp
					prev[v] = &minPos
				}
			}
		}
	}

	return CalcShortestFrom(from, to, prev)
}

func FindShortestPath(from []Pos, to Pos, vertices []Vertex) []Pos {
	fastAccess := make(map[Pos][]Pos, 0)
	for _, v := range vertices {
		fa, e := fastAccess[v.From]
		if !e {
			fastAccess[v.From] = []Pos{v.to}
		} else {
			fastAccess[v.From] = append(fa, v.to)
		}
	}

	dist := make(map[Pos]int, 0)
	prev := make(map[Pos]*Pos, 0)
	queueMap := make(map[Pos]bool, len(vertices))
	for _, v := range vertices {
		if slices.Contains(from, v.From) {
			dist[v.From] = 0
		} else {
			dist[v.From] = math.MaxInt
		}
		prev[v.From] = nil
		queueMap[v.From] = true
		dist[v.to] = math.MaxInt
		prev[v.to] = nil
		queueMap[v.to] = true
	}
	queue := make([]Pos, 0, len(queueMap))
	for p := range queueMap {
		queue = append(queue, p)
	}

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return dist[queue[i]] < dist[queue[j]]
		})
		minPos := queue[0]
		queue = queue[1:]
		for _, v := range fastAccess[minPos] {
			if slices.Contains(queue, v) {
				temp := dist[minPos] + 1
				if temp < dist[v] {
					dist[v] = temp
					prev[v] = &minPos
				}
			}
		}
	}
	shortest := make([]Pos, 0)
	for _, p := range from {
		if i := CalcShortestFrom(p, to, prev); len(shortest) == 0 || len(i) < len(shortest) {
			shortest = append(shortest, i...)
		}
	}

	return shortest
}

func CalcShortestFrom(from, to Pos, prev map[Pos]*Pos) []Pos {
	path := make([]Pos, 0)
	u := to
	if prev[to] != nil && to != from {
		for prev[u] != nil {
			path = append([]Pos{u}, path...)
			u = *prev[u]
		}
	}
	return path
}

func findPossibleMoves(pos Pos, world map[Pos]int) []Vertex {
	moves := make([]Vertex, 0)
	for _, v := range directions {
		possible, newPos := DirPossible(pos, v, world)
		if possible {
			moves = append(moves, Vertex{pos, newPos})
		}
	}
	return moves
}

type Vertex struct {
	From, to Pos
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}
func DirPossible(pos Pos, dir dir, world map[Pos]int) (bool, Pos) {
	var target Pos
	switch dir {
	case Up:
		target = Pos{pos.x, pos.y + 1}
	case Down:
		target = Pos{pos.x, pos.y - 1}
	case Left:
		target = Pos{pos.x - 1, pos.y}
	case Right:
		target = Pos{pos.x + 1, pos.y}
	}
	v, exists := world[target]
	if !exists || world[pos] < v-1 {
		return false, pos
	}
	return true, target
}

type tile int

const (
	Standard tile = iota
	Start
	End
)

type dir int

const (
	Up dir = iota
	Down
	Left
	Right
)

var directions = []dir{
	Up, Down, Left, Right,
}

func valueOf(r rune) (int, tile) {
	switch {
	case r == 'S':
		return int('z'-'a') + 1, Start
	case r == 'E':
		return int('z'-'a') + 1, End
	default:
		return int(r-'a') + 1, Standard
	}
}
func signOf(in int) string {
	return string(rune(in - 1 + 'a'))
}

type Pos struct {
	x, y int
}

type Dimensions Pos
