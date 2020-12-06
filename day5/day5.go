package main

import (
	"adventofcode/fileinput"
	"bufio"
	"io"
	"log"
	"math"
)

func main() {
	file, err := fileinput.OpenInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seats, err := parseSeats(file)
	if err != nil {
		log.Fatal(err)
	}

	missing := findMissingSeats(seats)
	for _, seat := range missing {
		log.Printf("Missing seat: %d", seat.ID())
	}
}

const height = 128
const width = 8

func findMissingSeats(seats []seat) []seat {
	var missing []seat

	present := map[seat]bool{}
	for _, seat := range seats {
		present[seat] = true
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			seat := newSeat(y, x)
			if _, ok := present[seat]; !ok {
				missing = append(missing, seat)
			}
		}
	}

	return missing
}

func parseSeats(input io.Reader) ([]seat, error) {
	var seats []seat

	scan := bufio.NewScanner(input)
	for scan.Scan() {
		line := scan.Text()
		seat := newSeat(findSeatPosition(line, height, width))
		seats = append(seats, seat)
	}

	return seats, scan.Err()
}

func findSeatPosition(seat string, height int, width int) (int, int) {
	a := int(math.Log2(float64(height)))
	b := int(math.Log2(float64(width)))

	y := 0
	x := 0

	for i := 0; i < a; i++ {
		if seat[i] == 'B' {
			y += 1 << (a - i - 1)
		}
	}

	for i := 0; i < b; i++ {
		if seat[a+i] == 'R' {
			x += 1 << (b - i - 1)
		}
	}

	return y, x
}

type seat struct {
	y int
	x int
}

func newSeat(y, x int) seat {
	return seat{y: y, x: x}
}

func (s seat) ID() int {
	return s.y*8 + s.x
}
