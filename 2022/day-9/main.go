package main

import (
	"aoc/day-9/plank"
	"fmt"
)

func main() {
	plank.Test()
	lines := plank.InputLineByLine("input.txt")
	moves, _ := plank.Parse(lines)
	field := plank.Initial10KnotField()
	for _, v := range moves {
		v.ApplyTo(field)
	}

	fmt.Println(field.VisitedByTCount())
}
