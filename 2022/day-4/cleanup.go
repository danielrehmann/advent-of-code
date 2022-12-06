package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type elfCleaningPlan struct {
	begin int
	end   int
}
type cleanupTeam struct {
	first  elfCleaningPlan
	second elfCleaningPlan
}

func (ct *cleanupTeam) hasRedundancy() bool {
	if ct.first.begin >= ct.second.begin && ct.first.end <= ct.second.end {
		return true
	}
	if ct.second.begin >= ct.first.begin && ct.second.end <= ct.first.end {
		return true
	}
	return false
}

func (ct *cleanupTeam) hasOverlap() bool {
	if ct.second.begin <= ct.first.begin && ct.first.begin <= ct.second.end {
		return true
	}
	if ct.first.begin <= ct.second.begin && ct.second.begin <= ct.first.end {
		return true
	}
	if ct.second.begin <= ct.first.end && ct.first.end <= ct.second.end {
		return true
	}
	if ct.first.begin <= ct.second.end && ct.second.end <= ct.first.end {
		return true
	}
	return false
}
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

func main() {
	lines := InputLineByLine("input.txt")
	plans := make([]cleanupTeam, 0, 100)
	for _, line := range lines {
		elves := strings.Split(line, ",")

		for i := 0; i < len(elves)-1; i += 2 {
			elfInput1 := strings.Split(elves[i], "-")
			elfInput2 := strings.Split(elves[i+1], "-")

			elf1b, _ := strconv.Atoi(elfInput1[0])
			elf1e, _ := strconv.Atoi(elfInput1[1])
			elf2b, _ := strconv.Atoi(elfInput2[0])
			elf2e, _ := strconv.Atoi(elfInput2[1])
			team := cleanupTeam{
				first:  elfCleaningPlan{begin: elf1b, end: elf1e},
				second: elfCleaningPlan{begin: elf2b, end: elf2e},
			}
			plans = append(plans, team)
		}
	}
	sumRedundant, sumOverlap := 0, 0
	for _, plan := range plans {
		if plan.hasRedundancy() {
			fmt.Println(plan)
			sumRedundant += 1
		}
		if plan.hasOverlap() {
			fmt.Println(plan)
			sumOverlap += 1
		}
	}
	fmt.Println(sumRedundant, sumOverlap)
}
