package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rucksackContent := parseInputToRucksack1(input)

	sum := 0
	for _, v := range rucksackContent {
		duplicateValue := v.valueOfDuplicate()
		sum = sum + duplicateValue
	}

	input, err = os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	parseInputToRucksack2(input)

}

func parseInputToRucksack2(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var elfGroups [][]string
	var elves []string
	var badges []Item
	for scanner.Scan() {
		if text := scanner.Text(); text != "" {
			elves = append(elves, text)
			if len(elves)%3 == 0 {
				fmt.Println(elves)
				badges = append(badges, getBadge(elves))
				elfGroups = append(elfGroups, elves)
				elves = nil
			}
		}
	}
	sum := 0
	for _, v := range badges {
		sum += v.valueOfRune()
	}
	fmt.Println(sum)
	return elfGroups
}

func getBadge(elves []string) Item {
	for _, v := range elves[0] {
		for _, x := range elves[1] {
			for _, y := range elves[2] {
				if x == v && v == y {
					fmt.Println(x, y, v)
					return Item{x}
				}
			}
		}
	}
	return Item{-1}
}

type rucksack struct {
	compartmentA []rune
	compartmentB []rune
}

func (r *rucksack) findDuplicate() (Item, string) {
	for _, a := range r.compartmentA {
		for _, b := range r.compartmentB {
			if a == b {
				return Item{a}, string(a)
			}
		}
	}
	return Item{rune(-1)}, ""
}

type Item struct {
	r rune
}

func (r *rucksack) valueOfDuplicate() int {
	v, _ := r.findDuplicate()
	return v.valueOfRune()

}

func (v *Item) valueOfRune() int {
	if unicode.IsLower(v.r) {
		return int(v.r-'a') + 1
	} else {
		return int(v.r-'A') + 27
	}
}

func parseInputToRucksack1(file *os.File) []rucksack {
	scanner := bufio.NewScanner(file)
	var rucksacks []rucksack
	for scanner.Scan() {
		if text := scanner.Text(); text != "" {
			compartmentSize := len(text) / 2
			var compartmentA []rune
			var compartmentB []rune
			for i, v := range text {
				if i < compartmentSize {
					compartmentA = append(compartmentA, v)
				} else {
					compartmentB = append(compartmentB, v)
				}
			}
			for i, j := 0, len(compartmentB)-1; i < j; i, j = i+1, j-1 {
				compartmentB[i], compartmentB[j] = compartmentB[j], compartmentB[i]
			}
			rucksacks = append(rucksacks, rucksack{compartmentA: compartmentA, compartmentB: compartmentB})
		}
	}
	return rucksacks

}
