package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Calculationg most Calories")
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	var elves []int
	currentCal := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentCal += calories
		} else {
			elves = append(elves, currentCal)
			currentCal = 0
		}
	}
	sort.Slice(elves, func(i, j int) bool { return elves[i] > elves[j] })

	sum := 0
	for _, value := range elves[:3] {
		sum += value
	}
	fmt.Println(sum)
}
