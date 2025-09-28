package main

import (
	"fmt"
	"os"
	"strings"
	"slices"
)

type Point struct {
	X int
	Y int
}

type PointList []Point

type Challenge struct {
	width int
	height int
	antennas map[rune]PointList
}

const chal_dist_sq int = 4

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input file")
		os.Exit(1)
	}

	chal := Parse(os.Args[1])

	var part1nodes PointList

	for freq := range chal.antennas {
		for _, ant1 := range chal.antennas[freq] {
			for _, ant2 := range chal.antennas[freq] {
				if ant1 == ant2 {
					continue
				}

				diff := ant1.Sub(ant2)

				node := ant1.Add(diff)
				dist1 := ant1.DistSquared(node)
				dist2 := ant2.DistSquared(node)

				if chal.InBounds(node) && (dist1 * chal_dist_sq == dist2 || dist2 * chal_dist_sq == dist1) && !slices.Contains(part1nodes, node) {
					part1nodes = append(part1nodes, node)
				}

				node = ant2.Sub(diff)
				dist1 = ant1.DistSquared(node)
				dist2 = ant2.DistSquared(node)

				if chal.InBounds(node) && (dist1 * chal_dist_sq == dist2 || dist2 * chal_dist_sq == dist1) && !slices.Contains(part1nodes, node) {
					part1nodes = append(part1nodes, node)
				}
			}
		}
	}

	var part2nodes PointList

	for freq := range chal.antennas {
		for _, ant1 := range chal.antennas[freq] {
			for _, ant2 := range chal.antennas[freq] {
				if ant1 == ant2 {
					continue
				}

				diff := ant1.Sub(ant2)

				node := ant1

				for chal.InBounds(node) {
					if !slices.Contains(part2nodes, node) {
						part2nodes = append(part2nodes, node)
					}
					node = node.Add(diff)
				}

				node = ant1

				for chal.InBounds(node) {
					if !slices.Contains(part2nodes, node) {
						part2nodes = append(part2nodes, node)
					}
					node = node.Sub(diff)
				}
			}
		}
	}

	fmt.Println("Part1:", len(part1nodes))
	fmt.Println("Part2:", len(part2nodes))
}

func Parse(filename string) Challenge {
	var chal Challenge
	chal.antennas = make(map[rune]PointList)

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))
	lines := strings.Split(text, "\n")

	chal.height = len(lines)
	chal.width  = len(lines[0])

	for y, line := range lines {
		if len(line) != chal.width {
			panic(fmt.Errorf("mismatched line length at line %d\n", y))
		}

		for x, char := range line {
			if char != '.' {
				chal.antennas[char] = append(chal.antennas[char], Point{X: x, Y: y})
			}
		}
	}
	return chal
}

func (c Challenge) InBounds(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < c.width && p.Y < c.height
}

func (a Point) Add(b Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a Point) Sub(b Point) Point {
	return Point{X: a.X - b.X, Y: a.Y - b.Y}
}

func (a Point) DistSquared(b Point) int {
	x := a.X - b.X
	y := a.Y - b.Y
	return x*x + y*y
}
