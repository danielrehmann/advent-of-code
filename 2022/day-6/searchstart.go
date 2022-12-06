package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getNextSliceOfRunes(scanner *bufio.Scanner) []rune {
	if scanner.Scan() {
		return []rune(scanner.Text())
	}
	return make([]rune, 0)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		panic("Could not open file input.txt")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	runes := getNextSliceOfRunes(scanner)
	currentStart := -1
	startTokenAfter := 14
	for i := 0; i < len(runes)-startTokenAfter+1; i++ {
		distinct := make([]rune, 0, startTokenAfter)
		for j := 0; j < startTokenAfter; j++ {
			if contains(&distinct, runes[i+j]) {
				continue
			} else {
				distinct = append(distinct, runes[i+j])
			}
		}
		if len(distinct) == startTokenAfter {
			currentStart = i + startTokenAfter
			break
		}
	}
	fmt.Println(currentStart)
}

func contains(distinct *[]rune, r rune) bool {
	for _, v := range *distinct {
		if v == r {
			return true
		}
	}
	return false
}
