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
		total += group.size()
	}

	fmt.Printf("Total: %d\n", total)
}

type question int16
type group map[question]bool

func newGroup() group {
	return map[question]bool{}
}

func (g *group) mark(q question) {
	(*g)[q] = true
}

func (g group) size() int {
	return len(g)
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
		for i := 0; i < len(line); i++ {
			q := question(line[i])
			g.mark(q)
		}
	}

	return groups, scanner.Err()
}
