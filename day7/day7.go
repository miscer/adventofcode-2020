package main

import (
	"adventofcode/day7/bags"
	"adventofcode/day7/parser"
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

	bags, err := parseBags(file)
	count := getParentBags(bags, "shiny gold")
	fmt.Println(count)
}

type bagMap map[string]bags.Bag

func parseBags(input io.Reader) (bagMap, error) {
	bags := bagMap{}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		bag, err := parser.ParseBag(line)
		if err != nil {
			return bags, err
		}
		bags[bag.Color] = bag
	}

	return bags, scanner.Err()
}

func getParentBags(bags bagMap, color string) (count int) {
	parents := map[string]bool{}

	var add func(color string)
	add = func(color string) {
		for _, bag := range bags {
			if bag.Contains(color) {
				parents[bag.Color] = true
				add(bag.Color)
			}
		}
	}
	add(color)

	return len(parents)
}
