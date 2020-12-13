package main

import (
	"adventofcode/fileinput"
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
)

func main() {
	file, err := fileinput.OpenInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ratings, err := readRatings(file)
	chain := createChain(ratings)
	parts := findParts(chain)

	result := 1
	for _, part := range parts {
		result *= countArrangements(part, 0)
	}

	log.Printf("result: %d", result)
}

func readRatings(reader io.Reader) (ratings []int, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		var rating int
		rating, err = strconv.Atoi(line)
		if err != nil {
			return
		}

		ratings = append(ratings, rating)
	}
	err = scanner.Err()
	return
}

func createChain(ratings []int) chain {
	chain := make([]int, len(ratings)+2)
	last := len(chain) - 1

	copy(chain[1:last], ratings)
	sort.Ints(chain[1:last])

	chain[0] = 0
	chain[last] = chain[last-1] + 3

	return chain
}

func findParts(ch chain) []chain {
	var parts []chain
	var current chain

	for i := 0; i < len(ch)-1; i++ {
		current = append(current, ch[i])

		if ch[i+1]-ch[i] >= 3 {
			parts = append(parts, current)
			current = nil
		}
	}

	current = append(current, ch[len(ch)-1])
	parts = append(parts, current)

	return parts
}

func countArrangements(ch chain, start int) int {
	total := 1

	for i := start; i < len(ch)-2; i++ {
		if ch[i+2]-ch[i] <= 3 {
			next := ch.remove(i + 1)
			total += countArrangements(next, i)
		}
	}

	return total
}

type chain []int

func (ch chain) remove(i int) chain {
	next := make([]int, len(ch)-1)
	copy(next[:i], ch[:i])
	copy(next[i:], ch[i+1:])
	return next
}
