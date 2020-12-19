package main

import (
	"reflect"
	"testing"
)

func TestDirection_Delta(t *testing.T) {
	tests := []struct {
		name  string
		angle int
		x     int
		y     int
	}{
		{name: "north", angle: 0, x: 0, y: 1},
		{name: "east", angle: 90, x: 1, y: 0},
		{name: "south", angle: 180, x: 0, y: -1},
		{name: "west", angle: 270, x: -1, y: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Direction{
				angle: tt.angle,
			}
			gotx, goty := d.Delta()
			if gotx != tt.x {
				t.Errorf("Delta() x = %v, want %v", gotx, tt.x)
			}
			if goty != tt.y {
				t.Errorf("Delta() y = %v, want %v", goty, tt.y)
			}
		})
	}
}

func TestVector_Turn(t *testing.T) {
	tests := []struct {
		name   string
		vector Vector
		angle  int
		want   Vector
	}{
		{name: "left", vector: Vector{1, 1}, angle: -90, want: Vector{-1, 1}},
		{name: "right", vector: Vector{1, 1}, angle: 90, want: Vector{1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.vector
			v.Turn(tt.angle)

			if !reflect.DeepEqual(v, tt.want) {
				t.Errorf("want: %s, got: %s", tt.want, v)
			}
		})
	}
}
