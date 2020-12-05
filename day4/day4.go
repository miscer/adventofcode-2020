package main

import (
	"adventofcode/fileinput"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := fileinput.OpenInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passes, err := readEntries(file)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, pass := range passes {
		if pass.IsValid() {
			count++
		}
	}

	fmt.Println(count)
}

type passport struct {
	birthYear      *string
	issueYear      *string
	expirationYear *string
	height         *string
	hairColor      *string
	eyeColor       *string
	passportID     *string
	countryID      *string
}

func (p passport) IsValid() bool {
	return p.birthYear != nil && p.issueYear != nil && p.expirationYear != nil &&
		p.height != nil && p.hairColor != nil && p.eyeColor != nil && p.passportID != nil
}

func readEntries(file *os.File) ([]passport, error) {
	passes := []passport{{}}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()

		if line == "" {
			passes = append(passes, passport{})
		}

		pass := &passes[len(passes)-1]
		err := readLine(line, pass)
		if err != nil {
			return passes, err
		}
	}

	if err := scan.Err(); err != nil {
		return passes, err
	}

	return passes, nil
}

func readLine(line string, p *passport) error {
	scan := bufio.NewScanner(strings.NewReader(line))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		word := scan.Text()
		parts := strings.Split(word, ":")
		key, value := parts[0], parts[1]

		switch key {
		case "byr":
			p.birthYear = &value
		case "iyr":
			p.issueYear = &value
		case "eyr":
			p.expirationYear = &value
		case "hgt":
			p.height = &value
		case "hcl":
			p.hairColor = &value
		case "ecl":
			p.eyeColor = &value
		case "pid":
			p.passportID = &value
		case "cid":
			p.countryID = &value
		}
	}

	return scan.Err()
}
