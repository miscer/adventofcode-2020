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

	ship := Ship{waypoint: Vector{x: 10, y: 1}}
	for _, i := range actions {
		ship = i.Move(ship)
		log.Printf("pos: %s, wp: %s", ship.Position(), ship.Waypoint())
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
	position Vector
	waypoint Vector
}

func (s *Ship) Position() *Vector {
	return &s.position
}

func (s *Ship) Waypoint() *Vector {
	return &s.waypoint
}

func (s *Ship) Distance() int {
	x, y := s.position.Get()
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
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
	ship.Waypoint().Add(m.distance*dx, m.distance*dy)
	return ship
}

type Forward struct {
	distance int
}

func (f Forward) Move(ship Ship) Ship {
	dx, dy := ship.Waypoint().Get()
	ship.Position().Add(f.distance*dx, f.distance*dy)
	return ship
}

type Turn struct {
	angle int
}

func (t Turn) Move(ship Ship) Ship {
	ship.Waypoint().Turn(t.angle)
	return ship
}

type Vector struct{ x, y int }

func (v Vector) Get() (int, int) {
	return v.x, v.y
}

func (v *Vector) Set(x, y int) {
	v.x = x
	v.y = y
}

func (v *Vector) Add(x, y int) {
	v.x += x
	v.y += y
}

func (v *Vector) Turn(deg int) {
	rad := float64(deg) / 180 * math.Pi * -1

	x := float64(v.x)*math.Cos(rad) - float64(v.y)*math.Sin(rad)
	y := float64(v.x)*math.Sin(rad) + float64(v.y)*math.Cos(rad)

	v.x, v.y = int(math.Round(x)), int(math.Round(y))
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d;%d)", v.x, v.y)
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
