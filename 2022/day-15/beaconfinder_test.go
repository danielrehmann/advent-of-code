package day15

import (
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
			got := LoadFromFile(tt.args.fn, tt.args.row)
			//fmt.Println(got.String())
			if bir := BlockedPositionsInRow(&got, tt.args.row); bir != tt.want {
				t.Fatalf("BlockedPositionsInRow() = %d, want %d", bir, tt.want)
			}

		})
	}
}
