package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

type move struct {
	amount int
	from   int
	to     int
}

func (m *move) applyTo(s *[][]string) {
	fmt.Println(s)
	from := (*s)[m.from]
	to := (*s)[m.to]
	to = append(to, from[len(from)-m.amount:]...)
	from = from[:len(from)-m.amount]
	(*s)[m.from] = from
	(*s)[m.to] = to
}

//func (m *move) applyToOld(s *[][]string) {
// 	fmt.Println(s)
// 	from := (*s)[m.from]
// 	to := (*s)[m.to]
// 	for i := 0; i < m.amount; i++ {
// 		to = append(to, from[len(from)-1])
// 		from = from[:len(from)-1]
// 	}
// 	(*s)[m.from] = from
// 	(*s)[m.to] = to
// }

func main() {
	lines := InputLineByLine("input.txt")
	str, _ := regexp.Compile(`.([A-Z\s])...([A-Z\s])...([A-Z\s])...([A-Z\s])...([A-Z\s])...([A-Z\s])...([A-Z\s])...([A-Z\s])...([A-Z\s]).`)
	mr, _ := regexp.Compile(`^move (\d+) from (\d+) to (\d+)$`)
	initialLines := [][]string{make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0)}
	moves := make([]move, 0)
	for _, v := range lines {
		if res := str.FindStringSubmatch(v); res != nil {
			for i, r := range res[1:] {
				if r != " " {
					initialLines[i] = append([]string{r}, initialLines[i]...)
				}
			}
		}
		if res := mr.FindStringSubmatch(v); res != nil {
			from, _ := strconv.Atoi(res[2])
			to, _ := strconv.Atoi(res[3])
			count, _ := strconv.Atoi(res[1])
			moves = append(moves, move{count, from - 1, to - 1})
		}
	}
	fmt.Println(initialLines)
	for _, m := range moves {
		m.applyTo(&initialLines)
	}
	fmt.Println(initialLines)
	for _, l := range initialLines {
		fmt.Print(l[len(l)-1])
	}

}
