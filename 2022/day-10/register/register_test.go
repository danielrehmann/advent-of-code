package register

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Command
	}{
		{"Noop can be parsed", "noop", noop{}},
		{"Addx can be parsed", "addx 1", addx{1}},
		{"Addx minus can be parsed", "addx -1", addx{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommand(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addx_Steps(t *testing.T) {
	type fields struct {
		value []Command
	}
	type expected struct {
		value, cycles int
	}
	tests := []struct {
		name   string
		fields fields
		want   expected
	}{
		{"single add 1", fields{[]Command{addx{1}}}, expected{1, 2}},
		{"two add 1", fields{[]Command{addx{1}, addx{1}}}, expected{2, 4}},
		{"single add -1", fields{[]Command{addx{-1}}}, expected{-1, 2}},
		{"add 1 and noop", fields{[]Command{addx{1}, noop{}}}, expected{1, 3}},
		{"add 15 and -11", fields{[]Command{addx{15}, addx{-11}}}, expected{4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			steps := make([]func(int) int, 0)
			for i := 0; i < len(tt.fields.value); i++ {
				steps = append(steps, tt.fields.value[i].Steps()...)
			}
			cycle, value := 0, 0
			for _, v := range steps {
				value = v(value)
				cycle = cycle + 1
			}
			if value != tt.want.value || cycle != tt.want.cycles {
				t.Fatalf("Steps() = expected cycles=%d, value=%d but was %d, %d", tt.want.cycles, tt.want.value, cycle, value)
			}
		})
	}
}
