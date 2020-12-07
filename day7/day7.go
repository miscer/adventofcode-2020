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
	count := countChildBags(bags, "shiny gold")
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

func countChildBags(bags bagMap, color string) (count int) {
	bag := bags[color]

	for c, n := range bag.Contents {
		count += n * (countChildBags(bags, c) + 1)
	}

	return
}
