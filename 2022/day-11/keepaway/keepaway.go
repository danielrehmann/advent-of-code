package day11

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func LoadKeepaway(lines []string) ([]Monkey, [][]int, int) {
	monkeys := make([]Monkey, 0, 10)
	items := make([][]int, 0, 10)
	sum := 1
	for i := 6; i < len(lines)+1; i = i + 7 {
		m, i, div := NewMonkey(lines[i-6 : i])
		monkeys = append(monkeys, m)
		items = append(items, i)
		sum *= div
	}
	return monkeys, items, sum
}

func Round(monkeys []Monkey, items [][]int, worryManager func(int) int) ([][]int, []int) {
	interacted := make([]int, len(monkeys))
	for id, v := range monkeys {
		for _, i := range items[id] {
			ni := v.operation(i)
			nir := worryManager(ni)
			nm := v.NextMonkey(nir)
			interacted[id] = interacted[id] + 1
			items[nm] = append(items[nm], nir)
		}
		items[id] = items[id][:0]
	}
	return items, interacted
}

type nextMonkey func(testMe int) int
type operation func(in int) int

type Monkey struct {
	operation  operation
	NextMonkey nextMonkey
}

var monkeyItemsRegex, _ = regexp.Compile(`Starting items: (.*)$`)
var monkeyOperationRegex, _ = regexp.Compile(`Operation: new = old (.) (\S+)$`)
var monkeyTestRegex, _ = regexp.Compile(`Test: divisible by (\d+)`)
var monkeyTargetRegex, _ = regexp.Compile(`If \w+: throw to monkey (\d+)`)

func NewMonkey(s []string) (Monkey, []int, int) {
	itmesString := monkeyItemsRegex.FindStringSubmatch(s[1])[1]
	operationRes := monkeyOperationRegex.FindStringSubmatch(s[2])
	operatorString, operationChangeString := operationRes[1], operationRes[2]
	targetTrue, _ := strconv.Atoi(monkeyTargetRegex.FindStringSubmatch(s[4])[1])
	targetFalse, _ := strconv.Atoi(monkeyTargetRegex.FindStringSubmatch(s[5])[1])

	itemsStrSl := strings.Split(itmesString, ", ")
	items := make([]int, 0, len(itemsStrSl))
	for _, v := range itemsStrSl {
		item, _ := strconv.Atoi(v)
		items = append(items, item)
	}
	var oldOrValue func(old int) int
	switch operationChangeString {
	case "old":
		oldOrValue = func(old int) int { return old }
	default:
		value, _ := strconv.Atoi(operationChangeString)
		oldOrValue = func(old int) int { return value }
	}

	var ops operation
	switch operatorString {
	case "+":
		ops = func(in int) int {
			return in + oldOrValue(in)
		}
	case "*":
		ops = func(in int) int {
			return in * oldOrValue(in)
		}
	}
	testDivisibleBy, _ := strconv.Atoi(monkeyTestRegex.FindStringSubmatch(s[3])[1])
	nextMonkey := func(test int) int {
		if test%testDivisibleBy == 0 {
			return targetTrue
		} else {
			return targetFalse
		}
	}
	return Monkey{ops, nextMonkey}, items, testDivisibleBy
}
