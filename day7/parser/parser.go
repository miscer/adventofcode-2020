package parser

import (
	"adventofcode/day7/bags"
	"fmt"
	"regexp"
	"strconv"
)

func ParseBag(input string) (bags.Bag, error) {
	bag := bags.Bag{Contents: map[string]int{}}

	re := regexp.MustCompile("^([\\w ]+) bags contain ([\\w ,]+).$")
	match1 := re.FindStringSubmatch(input)
	if match1 == nil {
		return bag, fmt.Errorf("invalid input: %s", input)
	}
	bag.Color = match1[1]

	re = regexp.MustCompile("(\\d+) ([\\w ]+) bags?")
	match2 := re.FindAllStringSubmatch(match1[2], -1)
	for _, match3 := range match2 {
		color := match3[2]
		count, err := strconv.Atoi(match3[1])
		if err != nil {
			return bag, err
		}

		bag.Contents[color] = count
	}

	return bag, nil
}
