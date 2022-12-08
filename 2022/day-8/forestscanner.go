package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := LoadInput("input.txt")
	visible := make(map[key]bool)
	allVisible := make([]key, 0, input.Width()*input.Height())
	for i := 0; i < input.Width(); i++ {
		allVisible = append(allVisible, input.TreesVisibleInCol(i)...)
	}
	for i := 0; i < input.Height(); i++ {
		allVisible = append(allVisible, input.TreesVisibleInRow(i)...)
	}
	for _, v := range allVisible {
		visible[v] = true
	}
	fmt.Println(visible, len(visible))

	sceenicRating := make(map[key]int)
	keys := make([]key, 0, input.Height()*input.Width())
	for i := 0; i < input.Height(); i++ {
		for j := 0; j < input.Width(); j++ {
			keys = append(keys, key{i, j})
		}
	}
	for _, k := range keys {
		//k := key{2, 1}
		if k.col == 0 || k.row == 0 {
			continue
		}
		treesLeft := input.TreesLeft(k)
		treesRight := input.TreesRight(k)
		treesAbove := input.treesAbove(k)
		treesBelow := input.treesBelow(k)
		tree := input.Tree(k)
		s := calculateSceenicRating(tree, [][]int{treesRight, treesLeft, treesAbove, treesBelow})
		sceenicRating[k] = s
	}
	fmt.Println("final", sceenicRating)

	max := 0
	for k, v := range sceenicRating {
		if v >= max {
			max = v
			fmt.Println("new max", k, v)
		}
	}
}

// tried 278052 - wrong because reading hard... all trees need to be checked not the output from part1

func calculateSceenicRating(tree int, neighbors [][]int) int {
	value := []int{0, 0, 0, 0}
	for i, n := range neighbors {
		for _, v := range n {
			value[i] = value[i] + 1
			if v >= tree {
				break
			}
		}
	}
	return value[0] * value[1] * value[2] * value[3]
}

func (f *Forest) Tree(k key) int {
	fmt.Println("at", k, "val", f.trees[k.row][k.col])
	return f.trees[k.row][k.col]
}

func (f *Forest) TreesLeft(k key) []int {
	trees := make([]int, 0)
	if k.col-1 < 0 {
		return trees
	}
	treeInv := f.trees[k.row][0:k.col]
	for i, j := 0, len(treeInv)-1; i < j; i, j = i+1, j-1 {
		treeInv[i], treeInv[j] = treeInv[j], treeInv[i]
	}
	return treeInv
}
func (f *Forest) TreesRight(k key) []int {
	row := f.trees[k.row]
	if k.col+1 > f.Height()-1 {
		return make([]int, 0)
	}
	trees := row[k.col+1:]
	return trees
}

func (f *Forest) treesAbove(k key) []int {
	trees := make([]int, 0)
	if k.row == 0 {
		return trees
	}
	for i := k.row - 1; i >= 0; i-- {
		trees = append(trees, f.trees[i][k.col])
	}
	return trees
}
func (f *Forest) treesBelow(k key) []int {
	trees := make([]int, 0)
	for i := k.row + 1; i < f.Height(); i++ {
		trees = append(trees, f.trees[i][k.col])
	}
	return trees
}

func LoadInput(s string) Forest {
	file, _ := os.Open(s)
	scanner := bufio.NewScanner(file)
	rows := make([][]int, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		column := make([]int, 0, len(line))
		for _, v := range line {
			size, _ := strconv.Atoi(string(v))
			column = append(column, size)
		}
		rows = append(rows, column)
	}
	return Forest{rows}
}

type Forest struct {
	trees [][]int
}

func (f *Forest) Width() int {
	if len(f.trees) < 1 {
		return -1
	}
	return len(f.trees[0])
}
func (f *Forest) Height() int {
	return len(f.trees)
}

func (f *Forest) TreesVisibleInRow(r int) []key {
	row := f.trees[r]
	treeVisible := make([]key, 0, len(row))
	highestX, highestY := -1, -1
	for x, y := 0, len(row)-1; x < len(row); x, y = x+1, y-1 {
		treeX, treeY := row[x], row[y]
		if treeX > highestX {

			treeVisible = append(treeVisible, key{r, x})
			highestX = treeX
		}
		if treeY > highestY {

			treeVisible = append(treeVisible, key{r, y})
			highestY = treeY
		}
	}
	return treeVisible

}

func (f *Forest) TreesVisibleInCol(column int) []key {
	treeVisible := make([]key, len(f.trees))
	highestX, highestY := -1, -1
	for x, y := 0, len(f.trees)-1; x < len(f.trees); x, y = x+1, y-1 {
		treeX, treeY := f.trees[x][column], f.trees[y][column]
		if treeX > highestX {
			treeVisible = append(treeVisible, key{x, column})
			highestX = treeX
		}
		if treeY > highestY {

			treeVisible = append(treeVisible, key{y, column})

			highestY = treeY
		}
	}
	return treeVisible
}

type key struct {
	col, row int
}
