package register

import (
	"strconv"
	"strings"
)

type Command interface {
	Steps() []func(int) int
}

type noop struct {
}

func (n noop) Steps() []func(int) int {
	return []func(int) int{identity}
}
func identity(in int) int {
	return in
}

type addx struct {
	value int
}

func (n addx) Steps() []func(int) int {
	return []func(int) int{identity, add(n.value)}
}

func add(i int) func(int) int {
	return func(inner int) int {
		return inner + i
	}
}

func ParseCommand(input string) Command {
	c := strings.Split(input, " ")
	switch c[0] {
	case "noop":
		return noop{}
	case "addx":
		value, _ := strconv.Atoi(c[len(c)-1])
		return addx{value}
	}
	return noop{}
}
