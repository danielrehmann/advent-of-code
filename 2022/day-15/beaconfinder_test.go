package day15

import (
	"fmt"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	type args struct {
		fn  string
		row int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example.txt", args{"example.txt", 10}, 26},
		{"input.txt", args{"input.txt", 2000000}, 5511201},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadFromFile(tt.args.fn)
			//fmt.Println(got.String())
			if bir := BlockedPositionsInRow(&got, tt.args.row); bir != tt.want {
				t.Fatalf("BlockedPositionsInRow() = %d, want %d", bir, tt.want)
			}
		})
	}
}

func TestBlockedPositionsInRow(t *testing.T) {
	type args struct {
		g Grid
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"MiniTest", args{Grid{[]Scan{NewScan(Pos{0, 0}, Pos{0, 1})}}, 1}, 0},
		{"MiniTest", args{Grid{[]Scan{NewScan(Pos{0, 0}, Pos{0, 1})}}, 0}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlockedPositionsInRow(&tt.args.g, tt.args.y); got != tt.want {
				fmt.Println(tt.args.g.String())
				t.Errorf("BlockedPositionsInRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindFreePos(t *testing.T) {
	type args struct {
		fn         string
		upperLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example.txt", args{"example.txt", 20}, 56000011},
		{"input.txt", args{"input.txt", 4000000}, 11318723411840},
	}
	for _, tt := range tests {
		grid := LoadFromFile(tt.args.fn)

		t.Run(tt.name, func(t *testing.T) {
			if got := calc(FindFreePos(&grid, tt.args.upperLimit)); got != tt.want {
				t.Errorf("FindFreePos() = %d, want %d", got, tt.want)
			}
		})
	}
}

func calc(got Pos) int {
	return got.x*4000000 + got.y
}
