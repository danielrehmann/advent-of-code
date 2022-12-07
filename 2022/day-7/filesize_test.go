package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestControl(t *testing.T) {
	file, _ := os.Open("test.txt")
	system := Calculate(bufio.NewScanner(file))

	fmt.Println(system.baseDir.Size())
	fmt.Println(system.allDirsLessThan(100000))
}
