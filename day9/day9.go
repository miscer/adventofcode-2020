package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
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
	if !found {
		log.Fatal("no invalid number found")
		return
	}

	set := findSet(nums, invalid)
	if set == nil {
		log.Fatal("no set of numbers found")
	}

	min, max := findMinMax(set)
	log.Printf("result: %d", min+max)
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

func findSet(nums []int, sum int) []int {
	for i := 0; i < len(nums)-1; i++ {
		acc := nums[i]
		set := []int{acc}

		for j := i + 1; j < len(nums); j++ {
			set = append(set, nums[j])
			acc += nums[j]

			if acc == sum {
				return set
			} else if acc > sum {
				break
			}
		}
	}

	return nil
}

func findMinMax(nums []int) (int, int) {
	min := math.MaxInt64
	max := 0

	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return min, max
}
