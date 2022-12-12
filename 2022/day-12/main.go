package main

import (
	day12 "aoc/day-12/wayfinder"
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	start, end, world, dim := day12.ParseMap(lines)
	fmt.Println(world, dim)
	av := day12.FindAllVertices(world, dim)
	path := day12.FindPath(start, end, av)
	fmt.Println(path, len(path))

	start, end, world, dim = day12.ParseMap(lines)
	world[start] = 1
	av = day12.FindAllVertices(world, dim)
	possibleStarts := make([]day12.Pos, 0)
	for p, v := range world {
		if v == 1 {
			possibleStarts = append(possibleStarts, p)
		}
	}
	sp := day12.FindShortestPath(possibleStarts, end, av)
	fmt.Println(sp, len(sp))

}
