package day13_test

import (
	day13 "aoc/day-13"
	"testing"
)

func Test_CorrectResult(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected int
	}{
		{"example test case", "example.txt", 13},
		{"real test case", "input.txt", 5013},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if res := day13.ResultForFile(tt.filename); res != tt.expected {
				t.Fatalf("ResultForFile() = %d ; want %d for file %s", res, tt.expected, tt.filename)
			}
		})
	}
}

func TestInRightOrder(t *testing.T) {
	type args struct {
		sp day13.SignalPair
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Pair 1", args{day13.NewSignalPair("[1,1,3,1,1]", "[1,1,5,1,1]")}, true},
		{"Pair 2", args{day13.NewSignalPair("[[1],[2,3,4]]", "[[1],4]")}, true},
		{"Pair 3", args{day13.NewSignalPair("[9]", "[[8,7,6]]")}, false},
		{"Pair 4", args{day13.NewSignalPair("[[4,4],4,4]", "[[4,4],4,4,4]")}, true},
		{"Pair 5", args{day13.NewSignalPair("[7,7,7,7]", "[7,7,7]")}, false},
		{"Pair 6", args{day13.NewSignalPair("[]", "[3]")}, true},
		{"Pair 7", args{day13.NewSignalPair("[[[]]]", "[[]]")}, false},
		{"Pair 8", args{day13.NewSignalPair("[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]")}, false},
		{"Pair 9", args{day13.NewSignalPair("[]", "[]")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day13.InRightOrder(tt.args.sp); got != tt.want {
				t.Errorf("InRightOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example test case", args{"example.txt"}, 140},
		{"real test case", args{"input.txt"}, 25038},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day13.OrderFile(tt.args.filename); got != tt.want {
				t.Errorf("OrderFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
