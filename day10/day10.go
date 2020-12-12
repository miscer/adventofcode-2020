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
	diffs := countDifferences(ratings)

	result := diffs[1] * diffs[3]
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

func countDifferences(ratings []int) map[int]int {
	chain := make([]int, len(ratings)+2)
	last := len(chain) - 1

	copy(chain[1:last], ratings)
	sort.Ints(chain[1:last])

	chain[0] = 0
	chain[last] = chain[last-1] + 3

	diffs := map[int]int{}
	for i := 0; i < len(chain)-1; i++ {
		diff := chain[i+1] - chain[i]
		diffs[diff]++
	}

	return diffs
}
