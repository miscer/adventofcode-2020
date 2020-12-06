package main

import (
	"adventofcode/fileinput"
	"bufio"
	"fmt"
	"io"
	"log"
)

func main() {
	file, err := fileinput.OpenInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	groups, err := parseGroups(file)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, group := range groups {
		total += group.count()
	}

	fmt.Printf("Total: %d\n", total)
}

type question int16
type group struct {
	answers map[question]int
	size    int
}

func newGroup() group {
	return group{
		answers: map[question]int{},
		size:    0,
	}
}

func (g *group) add() {
	g.size++
}

func (g *group) mark(q question) {
	g.answers[q]++
}

func (g group) count() int {
	count := 0

	for _, value := range g.answers {
		if value == g.size {
			count++
		}
	}

	return count
}

func parseGroups(reader io.Reader) ([]group, error) {
	groups := []group{newGroup()}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			groups = append(groups, newGroup())
			continue
		}

		g := &groups[len(groups)-1]
		g.add()

		for i := 0; i < len(line); i++ {
			q := question(line[i])
			g.mark(q)
		}
	}

	return groups, scanner.Err()
}
