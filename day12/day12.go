package main

import (
	"adventofcode/fileinput"
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
)

func main() {
	file, err := fileinput.OpenInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	actions, err := ParseActions(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(actions)

	ship := Ship{direction: East}
	for _, i := range actions {
		ship = i.Move(ship)

		x, y := ship.Position().Get()
		angle := ship.Direction().Angle()
		log.Printf("x: %d, y: %d, angle: %d", x, y, angle)
	}

	log.Printf("distance: %d", ship.Distance())
}

func ParseActions(reader io.Reader) (is []Instruction, err error) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		ch := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			return is, fmt.Errorf("invalid value: %s", line)
		}

		var instr Instruction
		switch ch {
		case 'N':
			instr = Move{direction: North, distance: value}
		case 'S':
			instr = Move{direction: South, distance: value}
		case 'E':
			instr = Move{direction: East, distance: value}
		case 'W':
			instr = Move{direction: West, distance: value}
		case 'F':
			instr = Forward{distance: value}
		case 'R':
			instr = Turn{angle: value}
		case 'L':
			instr = Turn{angle: -value}
		default:
			return is, fmt.Errorf("invalid instruction: %s", line)
		}

		is = append(is, instr)
	}

	return is, scanner.Err()
}

type Ship struct {
	direction Direction
	position  Position
}

func (s Ship) Position() Position {
	return s.position
}

func (s Ship) Direction() Direction {
	return s.direction
}

func (s Ship) Distance() int {
	x, y := s.position.Get()
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (s *Ship) Move(x, y int) {
	s.position.Move(x, y)
}

func (s *Ship) Turn(d Direction) {
	s.direction = d
}

type Instruction interface {
	Move(Ship) Ship
}

type Move struct {
	direction Direction
	distance  int
}

func (m Move) Move(ship Ship) Ship {
	dx, dy := m.direction.Delta()
	ship.Move(m.distance*dx, m.distance*dy)
	return ship
}

type Forward struct {
	distance int
}

func (f Forward) Move(ship Ship) Ship {
	dx, dy := ship.direction.Delta()
	ship.Move(f.distance*dx, f.distance*dy)
	return ship
}

type Turn struct {
	angle int
}

func (t Turn) Move(ship Ship) Ship {
	angle := (ship.Direction().Angle() + t.angle + 360) % 360
	ship.Turn(Direction{angle: angle})
	return ship
}

type Position struct{ x, y int }

func (p Position) Get() (int, int) {
	return p.x, p.y
}

func (p *Position) Move(x, y int) {
	p.x += x
	p.y += y
}

type Direction struct{ angle int }

func (d Direction) Angle() int {
	return d.angle
}

func (d Direction) Delta() (int, int) {
	angle := float64(d.angle) / 180 * math.Pi
	return int(math.Sin(angle)), int(math.Cos(angle))
}

var (
	North = Direction{angle: 0}
	South = Direction{angle: 180}
	East  = Direction{angle: 90}
	West  = Direction{angle: 270}
)
