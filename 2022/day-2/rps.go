package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type GameHint struct {
	opponent string
	answer   string
}

func (gh GameHint) key() string {
	return gh.opponent + " " + gh.answer
}

type Result struct {
	opponent int
	me       int
}

var results1 = map[string]Result{
	"A X": {1 + 3, 1 + 3},
	"A Y": {1 + 0, 2 + 6},
	"A Z": {1 + 6, 3 + 0},

	"B X": {2 + 6, 1 + 0},
	"B Y": {2 + 3, 2 + 3},
	"B Z": {2 + 0, 3 + 6},

	"C X": {3 + 0, 1 + 6},
	"C Y": {3 + 6, 2 + 0},
	"C Z": {3 + 3, 3 + 3},
}

var results = map[string]Result{
	"A X": {1 + 6, 3 + 0},
	"A Y": {1 + 3, 1 + 3},
	"A Z": {1 + 0, 2 + 6},

	"B X": {2 + 6, 1 + 0},
	"B Y": {2 + 3, 2 + 3},
	"B Z": {2 + 0, 3 + 6},

	"C X": {3 + 6, 2 + 0},
	"C Y": {3 + 3, 3 + 3},
	"C Z": {3 + 0, 1 + 6},
}

func (gh GameHint) resultFirstTest() int {
	return results1[gh.key()].me
}

func (gh GameHint) resultSecondTest() int {
	return results[gh.key()].me
}
func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	var lines []GameHint
	for scanner.Scan() {
		text := scanner.Text()
		temp := strings.Split(text, " ")
		lines = append(lines, GameHint{temp[0], temp[1]})
	}

	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	result1, result2 := 0, 0
	for _, v := range lines {
		result1 += v.resultFirstTest()
		result2 += v.resultSecondTest()
	}
	fmt.Println(result1, result2)
}
