package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	numbers, err := readNumbers(file)
	if err != nil {
		log.Fatal(err)
	}

	a, b, err := findMatchingEntries(numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found matching numbers: %d and %d\n", a, b)
}

func readNumbers(file *os.File) (numbers []int64, err error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid value in the input file: %s", line)
		}

		numbers = append(numbers, int64(number))
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return
}

func findMatchingEntries(numbers []int64) (int64, int64, error) {
	for i, a := range numbers {
		for _, b := range numbers[i+1:] {
			if a+b == 2020 {
				return a, b, nil
			}
		}
	}

	return 0, 0, errors.New("no matching entries found")
}
