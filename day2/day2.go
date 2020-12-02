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

	entries, err := readEntries(file)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, e := range entries {
		if e.isValid() {
			count++
		}
	}

	fmt.Printf("Found %d valid entries", count)
}

type entry struct {
	letter   string
	first    int
	second   int
	password string
}

func readEntries(file *os.File) (entries []entry, err error) {
	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()

		entry, err := parseEntry(line)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	if err = s.Err(); err != nil {
		return nil, err
	}

	return
}

func parseEntry(input string) (en entry, err error) {
	_, err = fmt.Sscanf(input, "%d-%d %1s: %s", &en.first, &en.second, &en.letter, &en.password)
	return
}

func (e *entry) isValid() bool {
	return (e.password[e.first-1:e.first] == e.letter) !=
		(e.password[e.second-1:e.second] == e.letter)
}
