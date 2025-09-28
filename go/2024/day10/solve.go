package main

import(
	"os"
	"fmt"
	"strings"
	_"strconv"
)

const (
	MAX_HEIGHT = 9
	MIN_HEIGHT = 0
)

type TopoMap [][]int

type Point struct {
	X int
	Y int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	topo := Parse(os.Args[1])
	peaks := topo.Peaks()
	score := 0

	for _, trailhead := range topo.Trailheads() {
		for _, peak := range peaks {
			if WalkTrailhead(trailhead, peak, topo) {
				score++
			}
		}
	}

	fmt.Println("Part1:", score)
}

func WalkTrailhead(head Point, peak Point, topo TopoMap) bool {
	if head == peak {
		return true
	}

	for _, neigh := range topo.Neighbors(head) {
		if topo.Height(neigh) == topo.Height(head) + 1 {
			if WalkTrailhead(neigh, peak, topo) {
				return true
			}
		}
	}

	return false
}


func Parse(filename string) TopoMap {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))

	var chal TopoMap

	lines := strings.Split(text, "\n")

	width := len(lines[0])

	for y, line := range lines {
		if len(line) != width {
			panic(fmt.Errorf("mismatched line len at line %d\n", y+1))
		}

		var mapline []int

		for _, char := range line {
			if char > '9' || char < '0' {
				panic(fmt.Errorf("char out of range at line %d: %s\n", y+1, string(char)))
			}
			mapline = append(mapline, int(char - '0'))
		}
		chal = append(chal, mapline)
	}

	return chal
}

func (m TopoMap) InBounds(p Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(m) && p.X < len(m[p.Y])
}

func (m TopoMap) Trailheads() []Point {
	var heads []Point

	for y := range m {
		for x := range m[y] {
			if m[y][x] == MIN_HEIGHT {
				heads = append(heads, Point{X: x, Y: y})
			}
		}
	}
	return heads
}

func (m TopoMap) Peaks() []Point {
	var peaks []Point

	for y := range m {
		for x := range m[y] {
			if m[y][x] == MAX_HEIGHT {
				peaks = append(peaks, Point{X: x, Y: y})
			}
		}
	}
	return peaks
}

func (m TopoMap) Height(p Point) int {
	if !m.InBounds(p) {
		panic(fmt.Errorf("point out of bounds %+v\n", p))
	}
	return m[p.Y][p.X]
}

func (m TopoMap) Neighbors(p Point) []Point {
	var neigh []Point
	if !m.InBounds(p) {
		panic(fmt.Errorf("point out of bounds %+v\n", p))
	}

	mod := []int{-1,0,1}

	for _, x := range mod {
		for _, y := range mod {
			if x == 0 && y == 0 {
				continue
			} else if x != 0 && y != 0 {
				continue
			}
			tmp := p
			tmp.X += x
			tmp.Y += y
			if m.InBounds(tmp) {
				neigh = append(neigh, tmp)
			}
		}
	}
	return neigh
}
