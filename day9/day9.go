package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	input, preamble, err := getArgs()
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	nums, err := readNumbers(input)
	invalid, found := findInvalidNumber(preamble, nums)

	if found {
		log.Printf("found invalid number: %d", invalid)
	} else {
		log.Print("no invalid number found")
	}
}

func getArgs() (io.ReadCloser, int, error) {
	preamble := flag.Int("preamble", 5, "size of the preamble")
	filename := flag.String("file", "", "input file")

	flag.Parse()

	var err error
	var input io.ReadCloser
	if *filename == "" {
		input = os.Stdin
	} else {
		input, err = os.Open(*filename)
		if err != nil {
			return nil, 0, fmt.Errorf("failed opening file: %w", err)
		}
	}

	return input, *preamble, nil
}

func readNumbers(reader io.Reader) (nums []int, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nums, err
		}
		nums = append(nums, num)
	}
	return nums, scanner.Err()
}

func findInvalidNumber(preamble int, nums []int) (int, bool) {
	for i := preamble; i < len(nums); i++ {
		num := nums[i]
		slice := nums[i-preamble : i]

		if !findPair(slice, num) {
			return num, true
		}
	}

	return 0, false
}

func findPair(nums []int, sum int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == sum {
				return true
			}
		}
	}
	return false
}
