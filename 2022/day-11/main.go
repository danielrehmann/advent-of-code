package main

import (
	day11 "aoc/day-11/keepaway"
	"fmt"
	"sort"
)

func identity(in int) int { return in }

func reduceBigNumber(cd int) func(int) int {
	return func(in int) int { return in % cd }
}

func main() {
	monkeys, items, commonDivisable := day11.LoadKeepaway(day11.InputLineByLine("input.txt"))
	interacted := make([]int, len(monkeys))
	for i := 0; i < 10000; i++ {
		newItems, interactedNew := day11.Round(monkeys, items, reduceBigNumber(commonDivisable))
		for i, v := range interactedNew {
			interacted[i] = interacted[i] + v
		}
		items = newItems
		//fmt.Printf("%d:\n %+v\n", i+1, interacted)
	}
	sort.Slice(interacted, func(i, j int) bool {
		return interacted[i] > interacted[j]
	})
	fmt.Println(interacted, interacted[0]*interacted[1])

}
