package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		log.Fatal("Input file name is missing")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed opening file: %v", err)
	}
	defer file.Close()

	m, err := readMap(file)
	if err != nil {
		log.Fatalf("Failed reading map: %v", err)
	}

	slopes := []struct{ dx, dy int }{
		{dx: 1, dy: 1},
		{dx: 3, dy: 1},
		{dx: 5, dy: 1},
		{dx: 7, dy: 1},
		{dx: 1, dy: 2},
	}

	result := 1
	for _, s := range slopes {
		c, err := countTrees(m, s.dx, s.dy)
		if err != nil {
			log.Fatalf("Failed counting trees: %v", err)
		}

		result *= c
	}

	fmt.Println(result)
}

type mapSquare int8

const (
	treeSquare mapSquare = iota
	emptySquare
)

func (s mapSquare) String() string {
	switch s {
	case treeSquare:
		return "🎄"
	case emptySquare:
		return "🌿"
	default:
		return "🤷"
	}
}

type worldMap struct {
	squares [][]mapSquare
}

func (m *worldMap) square(x int, y int) (mapSquare, error) {
	if y >= len(m.squares) || y < 0 {
		return emptySquare, fmt.Errorf("y coordinate out of range: %d", y)
	}

	row := m.squares[y]
	return row[x%len(row)], nil
}

func (m *worldMap) height() int {
	return len(m.squares)
}

func readMap(file *os.File) (m *worldMap, err error) {
	m = &worldMap{}
	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()

		var row []mapSquare
		for _, ch := range line {
			switch ch {
			case '#':
				row = append(row, treeSquare)
			case '.':
				row = append(row, emptySquare)
			}
		}

		m.squares = append(m.squares, row)
	}

	err = s.Err()

	return
}

func countTrees(m *worldMap, dx int, dy int) (int, error) {
	count := 0
	x := 0

	for y := 0; y < m.height(); y += dy {
		s, err := m.square(x, y)
		if err != nil {
			return 0, err
		}

		if s == treeSquare {
			count++
		}

		x += dx
	}

	return count, nil
}
