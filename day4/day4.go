package main

import (
	"adventofcode/fileinput"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	if p.birthYear == nil {
		return false
	}
	birthYear, err := strconv.Atoi(*p.birthYear)
	if err != nil || birthYear < 1920 || birthYear > 2002 {
		return false
	}

	if p.issueYear == nil {
		return false
	}
	issueYear, err := strconv.Atoi(*p.issueYear)
	if err != nil || issueYear < 2010 || issueYear > 2020 {
		return false
	}

	if p.expirationYear == nil {
		return false
	}
	expirationYear, err := strconv.Atoi(*p.expirationYear)
	if err != nil || expirationYear < 2020 || expirationYear > 2030 {
		return false
	}

	if p.height == nil {
		return false
	}
	heightRe := regexp.MustCompile("^(\\d+)(cm|in)$")
	heightMatches := heightRe.FindStringSubmatch(*p.height)
	if heightMatches == nil {
		return false
	}

	heightUnit := heightMatches[2]
	heightValue, err := strconv.Atoi(heightMatches[1])
	if err != nil {
		return false
	}
	if heightUnit == "cm" && (heightValue < 150 || heightValue > 193) {
		return false
	}
	if heightUnit == "in" && (heightValue < 59 || heightValue > 76) {
		return false
	}

	if p.hairColor == nil {
		return false
	}
	colorExpr := regexp.MustCompile("^#[a-f0-9]{6}$")
	if !colorExpr.MatchString(*p.hairColor) {
		return false
	}

	if p.eyeColor == nil {
		return false
	}
	eyeColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if _, ok := eyeColors[*p.eyeColor]; !ok {
		return false
	}

	if p.passportID == nil {
		return false
	}
	passIDRe := regexp.MustCompile("^\\d{9}$")
	if !passIDRe.MatchString(*p.passportID) {
		return false
	}

	return true
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
