package day14

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []Coordinate
	}{
		{"two neighboring nodes", args{"498,4 -> 498,5"}, []Coordinate{{498, 4}, {498, 5}}},
		{"two nodes with a space", args{"498,4 -> 498,6"}, []Coordinate{{498, 4}, {498, 5}, {498, 6}}},
		{"two nodes with a space horizontal", args{"498,4 -> 500,4"}, []Coordinate{{498, 4}, {499, 4}, {500, 4}}},
		{"two nodes with a space horizontal", args{"505,4 -> 500,4"}, []Coordinate{{505, 4}, {504, 4}, {503, 4}, {502, 4}, {501, 4}, {500, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, ParseLine(tt.args.s), tt.want)
		})
	}
}

//498,4 -> 498,6 -> 496,6
//503,4 -> 502,4 -> 502,9 -> 494,9

func TestParseInputFromFile(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name    string
		args    args
		want    World
		wantErr bool
	}{
		{"Example", args{"example.txt"}, World{
			{498, 4}: Rock, {498, 5}: Rock, {498, 6}: Rock, {497, 6}: Rock, {496, 6}: Rock,
			{503, 4}: Rock, {502, 4}: Rock, {502, 5}: Rock, {502, 6}: Rock, {502, 7}: Rock, {502, 8}: Rock, {502, 9}: Rock,
			{501, 9}: Rock, {500, 9}: Rock, {499, 9}: Rock, {498, 9}: Rock, {497, 9}: Rock, {496, 9}: Rock, {495, 9}: Rock, {494, 9}: Rock,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInputFromFile(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInputFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInputFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropNextSand(t *testing.T) {
	type args struct {
		world    World
		dropPos  Coordinate
		extremes ext
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"first drop", args{getFile("example.txt"), Coordinate{500, 0}, ExtremesOfWorld(getFile("example.txt"))}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropNextSand(tt.args.world, tt.args.dropPos, tt.args.extremes); got != tt.want {
				t.Errorf("DropNextSand() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.world.String())
		})
	}
}
func getFile(s string) World {
	exampleWorld, _ := ParseInputFromFile(s)
	return exampleWorld
}
func TestDropAllPossibleSand(t *testing.T) {
	type args struct {
		w           World
		startingPos Coordinate
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{getFile("example.txt"), Coordinate{500, 0}}, 24},
		{"input", args{getFile("input.txt"), Coordinate{500, 0}}, 1199},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropAllPossibleSand(tt.args.w, tt.args.startingPos); got != tt.want {
				t.Errorf("DropAllPossibleSand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropAllPossibleSandOnWorldWithFloor(t *testing.T) {
	type args struct {
		w           World
		startingPos Coordinate
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{getFile("example.txt"), Coordinate{500, 0}}, 93},
		{"input", args{getFile("input.txt"), Coordinate{500, 0}}, 23925},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DropAllPossibleSandWithFloor(tt.args.w, tt.args.startingPos); got != tt.want {
				t.Errorf("DropAllPossibleSandWithFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}
