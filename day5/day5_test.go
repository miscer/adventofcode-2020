package main

import "testing"

func Test_findSeatPosition(t *testing.T) {
	type args struct {
		seat   string
		height int
		width  int
	}
	tests := []struct {
		args args
		x    int
		y    int
	}{
		{args: args{seat: "FBFBBFFRLR", height: 128, width: 8}, y: 44, x: 5},
		{args: args{seat: "BFFFBBFRRR", height: 128, width: 8}, y: 70, x: 7},
		{args: args{seat: "FFFBBBFRRR", height: 128, width: 8}, y: 14, x: 7},
		{args: args{seat: "BBFFBBFRLL", height: 128, width: 8}, y: 102, x: 4},
	}

	for _, tt := range tests {
		t.Run(tt.args.seat, func(t *testing.T) {
			y, x := findSeatPosition(tt.args.seat, tt.args.height, tt.args.width)

			if y != tt.y {
				t.Errorf("findSeatPosition() y = %v, want %v", y, tt.y)
			}
			if x != tt.x {
				t.Errorf("findSeatPosition() x = %v, want %v", x, tt.x)
			}
		})
	}
}
