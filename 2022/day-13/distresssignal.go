package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ResultForFile(filename string) int {
	sum := 0
	for i, sp := range SignalPairsFromFile(filename) {
		fmt.Printf("== Pair %d ==\n", i+1)
		if InRightOrder(sp) {
			sum += (i + 1)
		}
		fmt.Println()
	}
	return sum
}

func OrderFile(filename string) int {
	signals := AllSignalsFromFile(filename)
	signals = append(signals, "[[2]]", "[[6]]")
	sort.Slice(signals, func(i, j int) bool {
		return InRightOrder(NewSignalPair(signals[i], signals[j]))
	})
	sum := 1
	for i, v := range signals {
		if v == "[[2]]" || v == "[[6]]" {
			sum *= i + 1
		}
	}
	return sum
}

func AllSignalsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		panic("Could not open file " + filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	signals := make([]string, 0, 100)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			signals = append(signals, t)
		}
	}
	return signals
}

func SignalPairsFromFile(filename string) []SignalPair {
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
	return ToComparePairs(lines)
}

type SignalPair [2]string

func ToComparePairs(lines []string) []SignalPair {
	pairs := make([]SignalPair, 0, len(lines)/3)
	for i := 2; i <= len(lines); i = i + 3 {
		pairs = append(pairs, NewSignalPair(lines[i-2], lines[i-1]))
	}
	return pairs
}

func NewSignalPair(left string, right string) SignalPair {
	return SignalPair{left, right}
}

func InRightOrder(sp SignalPair) bool {
	tokenizedLeft, tokenizedRight := tokenize(sp[0]), tokenize(sp[1])
	res, _ := compare(tokenizedLeft, tokenizedRight, 0)
	return res
}

func compare(tokenizedLeft, tokenizedRight []token, depth int) (bool, bool) {
	var dsBuilder strings.Builder
	for i := 0; i < depth; i++ {
		dsBuilder.WriteString("  ")
	}
	for i := 0; i < len(tokenizedLeft); i++ {
		left := tokenizedLeft[i]
		if len(tokenizedRight) <= i {
			fmt.Printf("%s- Right side ran out of items, so inputs are not in the right order\n", dsBuilder.String())
			return false, true
		}
		right := tokenizedRight[i]

		if left.isList() || right.isList() {
			fmt.Printf("%s- Compare %v vs %v\n", dsBuilder.String(), left.token(), right.token())
			res, fin := compare(left.token(), right.token(), depth+1)
			if fin {
				return res, true
			}
		} else {
			fmt.Printf("%s- Compare %d vs %d\n", dsBuilder.String(), left.toInt(), right.toInt())
			if left.toInt() > right.toInt() {
				fmt.Printf("%s- Right side is smaller, so inputs are not in the right order\n", dsBuilder.String())
				return false, true
			}
			if left.toInt() < right.toInt() {
				fmt.Printf("%s- Left side is smaller, so inputs are in the right order\n", dsBuilder.String())
				return true, true
			}
		}
	}
	if len(tokenizedLeft) < len(tokenizedRight) {
		fmt.Printf("%s- Left side ran out of items, so inputs are in the right order\n", dsBuilder.String())
		return true, true
	} else {
		return true, false
	}
}

func tokenize(str string) []token {
	if len(str) == 0 {
		return make([]token, 0)
	}
	switch {
	case strings.HasPrefix(str, "["):
		closingPos := -1
		openBraces := 1
		for i, v := range str[1:] {
			if v == '[' {
				openBraces = openBraces + 1
			}
			if v == ']' {
				openBraces = openBraces - 1
			}
			if openBraces == 0 {
				closingPos = i + 1
				break
			}
		}
		if closingPos == -1 {
			panic("no closing brace found in bracket")
		}
		inner := tokenizeInner(str[1:closingPos])

		remaining := make([]token, 0)
		if len(str) > closingPos+1 {
			remaining = tokenize(str[closingPos+1:])
		}
		return append([]token{inner}, remaining...)

	case str[0] == ',':
		return tokenize(str[1:])
	default:
		var next int
		nextcomma, nextInner := strings.Index(str, ","), strings.Index(str, "[")
		switch {
		case nextcomma == -1 && nextInner == -1:
			n, err := strconv.Atoi(str)
			if err != nil {
				panic("did not read number")
			}
			return []token{numberToken(n)}

		case nextcomma == -1 && nextInner > -1:
			next = nextInner
		case nextcomma > -1 && nextInner == -1:
			next = nextcomma
		case nextcomma >= nextInner:
			next = nextInner
		case nextcomma < nextInner:
			next = nextcomma
		}
		n, err := strconv.Atoi(str[:next])
		if err != nil {
			panic(fmt.Sprintf("did not read number %s, in %s", str[:next], str))
		}
		return append([]token{numberToken(n)}, tokenize(str[next:])...)

	}
}

func tokenizeInner(str string) token {
	return tokenlist(tokenize(str))
}

type tokenlist []token

func (il tokenlist) token() []token {
	return il
}
func (il tokenlist) isList() bool {
	return true
}
func (il tokenlist) toInt() int {
	if len(il) == 0 {
		return 0
	}
	return il[0].toInt()
}

type numberToken int

func (n numberToken) token() []token {
	return []token{n}
}
func (n numberToken) isList() bool {
	return false
}
func (n numberToken) toInt() int {
	return int(n)
}

type token interface {
	token() []token
	isList() bool
	toInt() int
}
