package main

import (
	"adventofcode/fileinput"
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	file, err := fileinput.OpenInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lo, err := readLayout(file)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Print(lo)
		fmt.Print("\n")

		var changes int
		lo, changes = step(lo)

		if changes == 0 {
			break
		}
	}

	fmt.Printf("Occupied: %d\n", lo.TotalOccupied())
}

func readLayout(reader io.Reader) (Layout, error) {
	layout := Layout{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		var row []Cell

		for _, char := range line {
			cell, err := ParseCell(char)
			if err != nil {
				return layout, err
			}

			row = append(row, cell)
		}

		layout.grid = append(layout.grid, row)
	}

	if err := scanner.Err(); err != nil {
		return layout, err
	}

	layout.width = len(layout.grid[0])
	layout.height = len(layout.grid)

	return layout, nil
}

func step(l Layout) (Layout, int) {
	next := NewLayout(l.width, l.height)
	changes := 0

	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			cell := l.Cell(x, y)
			occ := l.VisibleOccupied(x, y)

			if cell == EmptySeat && occ == 0 {
				next.Set(x, y, OccupiedSeat)
				changes++
			} else if cell == OccupiedSeat && occ >= 5 {
				next.Set(x, y, EmptySeat)
				changes++
			} else {
				next.Set(x, y, cell)
			}
		}
	}

	return next, changes
}

type Layout struct {
	width  int
	height int
	grid   [][]Cell
}

func NewLayout(width, height int) Layout {
	l := Layout{
		width:  width,
		height: height,
		grid:   make([][]Cell, height),
	}

	for i := range l.grid {
		l.grid[i] = make([]Cell, width)
	}

	return l
}

func (l Layout) String() string {
	var b strings.Builder

	for _, row := range l.grid {
		for _, cell := range row {
			b.WriteString(cell.String())
		}
		b.WriteString("\n")
	}

	return b.String()
}

func (l Layout) Cell(x int, y int) Cell {
	if y < 0 || y >= l.height {
		return OutOfBounds
	}
	if x < 0 || x >= l.width {
		return OutOfBounds
	}
	return l.grid[y][x]
}

func (l Layout) Set(x int, y int, cell Cell) {
	l.grid[y][x] = cell
}

func (l Layout) VisibleOccupied(x int, y int) int {
	occ := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			dx, dy := i, j

			for {
				cell := l.Cell(x+dx, y+dy)

				if cell == OutOfBounds || cell == EmptySeat {
					break
				}
				if cell == OccupiedSeat {
					occ++
					break
				}

				dx += i
				dy += j
			}
		}
	}

	return occ
}

func (l Layout) TotalOccupied() int {
	occ := 0

	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			if l.Cell(x, y) == OccupiedSeat {
				occ++
			}
		}
	}

	return occ
}

type Cell int

const (
	EmptySeat Cell = iota
	OccupiedSeat
	Floor
	OutOfBounds
)

func ParseCell(char int32) (Cell, error) {
	switch char {
	case 'L':
		return EmptySeat, nil
	case '.':
		return Floor, nil
	default:
		return Floor, fmt.Errorf("unexpected value: %c", char)
	}
}

func (c Cell) String() string {
	switch c {
	case EmptySeat:
		return "L"
	case OccupiedSeat:
		return "#"
	case Floor:
		return "."
	default:
		return "?"
	}
}
