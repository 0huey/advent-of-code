package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Right Direction = iota
	Left
)

type Rotation struct {
	Dir  Direction
	Dist int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	rot := Parse(os.Args[1])
	Part1(rot)
	Part2(rot)
}

func Parse(filename string) []Rotation {
	var rotations []Rotation

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))
	lines := strings.Split(text, "\n")

	for lineno, line := range lines {
		var rot Rotation

		line = strings.TrimSpace(line)

		if len(line) < 2 {
			panic(fmt.Errorf("invalid line %d %s", lineno, line))
		}

		switch line[0] {
		case 'R':
			rot.Dir = Right
		case 'L':
			rot.Dir = Left
		default:
			panic(fmt.Errorf("invalid direction at line %d %s", lineno, line))
		}

		rot.Dist, err = strconv.Atoi(line[1:])
		if err != nil {
			panic(fmt.Errorf("invalid number at line %d %s", lineno, line))
		}

		rotations = append(rotations, rot)
	}
	return rotations
}

func Part1(rotations []Rotation) {
	dial := 50
	zeros := 0

	for _, rot := range rotations {

		rot.Dist = rot.Dist % 100

		switch rot.Dir {
		case Right:
			dial += rot.Dist
		case Left:
			dial -= rot.Dist
		}

		if dial >= 100 {
			dial -= 100
		} else if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			zeros++
		}
	}
	fmt.Println("Part1:", zeros)
}

func Part2(rotations []Rotation) {
	dial := 50
	zeros := 0

	for _, rot := range rotations {
		switch rot.Dir {
		case Right:
			dial += rot.Dist
		case Left:
			dial -= rot.Dist
		}

		if dial == 0 {
			zeros++
		} else if dial < 0 {
			zeros += (dial*-1)/100 + 1
			dial = 100 - (dial*-1)%100
		} else {
			zeros += dial / 100
			dial = dial % 100
		}
	}
	fmt.Println("Part2:", zeros)
}
