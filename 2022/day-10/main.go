package main

import (
	"aoc/day-10/register"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	var commands []string
	for s.Scan() {
		commands = append(commands, s.Text())
	}
	var steps []func(int) int
	for _, v := range commands {
		c := register.ParseCommand(v)
		steps = append(steps, c.Steps()...)
	}
	value := 1
	interestedIn := make(map[int]int, 0)
	var crt strings.Builder
	allValues := make(map[int]int)

	allValues[1] = 1
	for i := 1; i <= len(steps); i++ {
		value = steps[i-1](value)
		allValues[i+1] = value
		if shiftedI := i + 21; shiftedI > 0 && shiftedI%40 == 0 {
			interestedIn[i+1] = value
		}
		fmt.Printf("In cycle %d value was %d\n", i, value)
	}
	fmt.Println(interestedIn)
	fmt.Println(allValues)

	sum := 0
	for k, v := range interestedIn {
		sum += k * v
	}
	for k := 1; k <= len(allValues); k++ {
		v := allValues[k]
		if i := (k - 1) % 40; -1 <= v-i && v-i <= 1 {
			crt.WriteString("#")
			fmt.Println(k, v, i, "#")
		} else {
			crt.WriteString(" ")
		}
		if k > 1 && (k)%40 == 0 {
			fmt.Println("newline")
			crt.WriteString("\n")
		}
	}

	println(sum)
	fmt.Printf(crt.String())

}
