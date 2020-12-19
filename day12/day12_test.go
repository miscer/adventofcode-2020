package main

import (
	"reflect"
	"testing"
)

func TestTurn_Move(t1 *testing.T) {
	tests := []struct {
		name  string
		angle int
		ship  Ship
		want  Ship
	}{
		{name: "right 90", angle: 90, ship: Ship{direction: North}, want: Ship{direction: East}},
		{name: "right 180", angle: 180, ship: Ship{direction: North}, want: Ship{direction: South}},
		{name: "right 270", angle: 270, ship: Ship{direction: North}, want: Ship{direction: West}},
		{name: "left 90", angle: -90, ship: Ship{direction: North}, want: Ship{direction: West}},
		{name: "left 180", angle: -180, ship: Ship{direction: North}, want: Ship{direction: South}},
		{name: "left 270", angle: -270, ship: Ship{direction: North}, want: Ship{direction: East}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Turn{
				angle: tt.angle,
			}
			if got := t.Move(tt.ship); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
