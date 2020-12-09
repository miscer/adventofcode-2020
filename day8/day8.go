package main

import (
	"adventofcode/fileinput"
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
)

func main() {
	file, err := fileinput.OpenInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	prog, err := parseProgram(file)
	if err != nil {
		log.Fatal(err)
	}

	acc, err := runProgram(prog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d", acc)
}

type instruction struct {
	operation string
	argument  int
}

type program []instruction

func parseProgram(input io.Reader) (program, error) {
	var prog program
	scanner := bufio.NewScanner(input)
	re := regexp.MustCompile("^(\\w+) ([+-]\\d+)$")

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		if matches == nil {
			return prog, fmt.Errorf("invalid line: %s", line)
		}

		op := matches[1]
		arg, err := strconv.Atoi(matches[2])
		if err != nil {
			return prog, fmt.Errorf("invalid argument: %s", matches[2])
		}

		instr := instruction{operation: op, argument: arg}
		prog = append(prog, instr)
	}

	return prog, scanner.Err()
}

type computer struct {
	prog program
	acc  int
	ip   int
}

func newComputer(prog program) computer {
	return computer{prog: prog}
}

func (c *computer) tick() error {
	if c.ip >= len(c.prog) {
		return fmt.Errorf("instruction pointer out of range: %d", c.ip)
	}

	instr := c.prog[c.ip]

	switch instr.operation {
	case "acc":
		c.acc += instr.argument
		c.ip++
	case "jmp":
		c.ip += instr.argument
	case "nop":
		c.ip++
	default:
		return fmt.Errorf("invalid operation: %s", instr.operation)
	}

	return nil
}

func runProgram(prog program) (int, error) {
	comp := newComputer(prog)
	ran := map[int]bool{}

	for {
		if _, ok := ran[comp.ip]; ok {
			break
		}

		ran[comp.ip] = true

		err := comp.tick()
		if err != nil {
			return 0, err
		}
	}

	return comp.acc, nil
}
