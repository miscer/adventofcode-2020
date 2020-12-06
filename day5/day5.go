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

	id, err := findHighestSeatID(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result: %d", id)
}

func findHighestSeatID(input io.Reader) (int, error) {
	var highest int

	scan := bufio.NewScanner(input)
	for scan.Scan() {
		line := scan.Text()
		y, x := findSeatPosition(line, 128, 8)
		id := (y * 8) + x

		if id > highest {
			highest = id
		}
	}

	return highest, scan.Err()
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
