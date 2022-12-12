package day12

import (
	"reflect"
	"testing"
)

func TestParseMap(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name  string
		args  args
		want  Pos
		want1 Pos
		want2 map[Pos]int
		want3 Dimensions
	}{
		{"smallMap", args{[]string{"Sy", "cE"}}, Pos{1, 0}, Pos{0, 1}, map[Pos]int{
			{0, 0}: 3,
			{0, 1}: 26,
			{1, 1}: 25,
			{1, 0}: 26,
		}, Dimensions{2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := ParseMap(tt.args.lines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseMap() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ParseMap() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("ParseMap() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func Test_findPossibleMoves(t *testing.T) {

	lines := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}
	_, _, world, _ := ParseMap(lines)
	type args struct {
		pos   Pos
		world map[Pos]int
	}
	tests := []struct {
		name string
		args args
		want []Vertex
	}{
		{"cant move", args{Pos{0, 0}, map[Pos]int{
			{0, 0}: 3,
			{0, 1}: 26,
			{1, 1}: 25,
			{1, 0}: 26,
		}}, make([]Vertex, 0)},
		{"moves", args{Pos{0, 1}, map[Pos]int{
			{0, 0}: 3,
			{0, 1}: 26,
			{1, 1}: 25,
			{1, 0}: 26,
		}}, []Vertex{
			{From: Pos{0, 1}, to: Pos{0, 0}},
			{From: Pos{0, 1}, to: Pos{1, 1}},
		}},
		{"with example", args{Pos{2, 0}, world}, []Vertex{
			{From: Pos{2, 0}, to: Pos{2, 1}},
			{From: Pos{2, 0}, to: Pos{1, 0}},
			{From: Pos{2, 0}, to: Pos{3, 0}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPossibleMoves(tt.args.pos, tt.args.world); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPossibleMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMovesToEnd(t *testing.T) {
	type args struct {
		world map[Pos]int
		dim   Dimensions
	}
	tests := []struct {
		name string
		args args
		want []Vertex
	}{
		{"vertices 2x1", args{map[Pos]int{
			{0, 0}: 25,
			{1, 0}: 26,
		}, Dimensions{2, 1}}, []Vertex{
			{From: Pos{0, 0}, to: Pos{1, 0}},
			{From: Pos{1, 0}, to: Pos{0, 0}},
		}},
		{"moves 2x1", args{map[Pos]int{
			{0, 0}: 24,
			{1, 0}: 26,
		}, Dimensions{2, 1}}, []Vertex{
			{From: Pos{1, 0}, to: Pos{0, 0}},
		}},
		{"moves 3x1", args{map[Pos]int{
			{0, 0}: 24,
			{1, 0}: 26,
			{2, 0}: 24,
		}, Dimensions{3, 1}}, []Vertex{
			{From: Pos{1, 0}, to: Pos{0, 0}},
			{From: Pos{1, 0}, to: Pos{2, 0}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAllVertices(tt.args.world, tt.args.dim); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMovesToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPath(t *testing.T) {
	type args struct {
		from     Pos
		to       Pos
		vertices []Vertex
	}
	tests := []struct {
		name string
		args args
		want []Pos
	}{
		{"vertices 2x1", args{
			Pos{0, 0}, Pos{3, 0}, []Vertex{
				{From: Pos{0, 0}, to: Pos{1, 0}},
				{From: Pos{1, 0}, to: Pos{2, 0}},
				{From: Pos{1, 0}, to: Pos{3, 0}},
				{From: Pos{2, 0}, to: Pos{3, 0}},
			}}, []Pos{{1, 0}, {3, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPath(tt.args.from, tt.args.to, tt.args.vertices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
