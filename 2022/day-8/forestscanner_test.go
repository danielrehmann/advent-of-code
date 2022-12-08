package main

import (
	"reflect"
	"testing"
)

func TestLoadInput(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Example can be loaded", args{"example.txt"}, [][]int{
				[]int{3, 0, 3, 7, 3},
				[]int{2, 5, 5, 1, 2},
				[]int{6, 5, 3, 3, 2},
				[]int{3, 3, 5, 4, 9},
				[]int{3, 5, 3, 9, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadInput(tt.args.s); !reflect.DeepEqual(got.trees, tt.want) {
				t.Errorf("LoadInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForest_Width(t *testing.T) {
	type fields struct {
		trees [][]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"example", fields{[][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Forest{
				trees: tt.fields.trees,
			}
			if got := f.Width(); got != tt.want {
				t.Errorf("Forest.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForest_Height(t *testing.T) {
	type fields struct {
		trees [][]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"example", fields{[][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Forest{
				trees: tt.fields.trees,
			}
			if got := f.Height(); got != tt.want {
				t.Errorf("Forest.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
