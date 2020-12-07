package main

import (
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
	fmt.Println(bags)
}

func parseBags(input io.Reader) ([]parser.ParsedBag, error) {
	var bags []parser.ParsedBag
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		bag, err := parser.ParseBag(line)
		if err != nil {
			return bags, err
		}
		bags = append(bags, bag)
	}

	return bags, scanner.Err()
}
