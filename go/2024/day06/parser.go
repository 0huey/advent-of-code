package main

import (
    "os"
    "strings"
    "fmt"
)

type Point struct {
    X int
    Y int
}

type AreaPoint byte
const (
    AreaBlank AreaPoint = iota
    AreaObstruction
    AreaStart
    AreaVisited
)

type AreaMap [][]AreaPoint

func Parse(filename string) (AreaMap, Point) {
    var area AreaMap
    var start Point = UnsetPoint()

    data, err := os.ReadFile(filename)
    if err != nil {
        panic(err)
    }

    str_data := strings.TrimSpace(string(data))
    lines := strings.Split(str_data, "\n")

    line_len := len(lines[0])

    for lineno := range lines {
        if len(lines[lineno]) != line_len {
            err = fmt.Errorf("mismatched line len at %d", lineno)
            panic(err)
        }

        var points []AreaPoint

        for i, p := range lines[lineno] {
            switch p {
                case '.':
                    points = append(points, AreaBlank)
                case '#':
                    points = append(points, AreaObstruction)
                case '^':
                    points = append(points, AreaStart)
                    if start != UnsetPoint() {
                        err = fmt.Errorf("multiple start points. First: %v, second: %d %d", start, i, lineno)
                        panic(err)
                    }
                    start = Point{X: i, Y: lineno}
                default:
                    err = fmt.Errorf("unknown text at %d %d", i, lineno)
                    panic(err)
            }
        }
        area = append(area, points)
    }
    return area, start
}

func PointInArea(area AreaMap, p Point) bool {
    return p.X >= 0 && p.Y >= 0 && p.Y < len(area) && p.X < len(area[0])
}

func ObjectAtPoint(area AreaMap, p Point) AreaPoint {
    return area[p.Y][p.X]
}
func SetObjectAtPoint(area AreaMap, p Point, v AreaPoint) {
    area[p.Y][p.X] = v
}

func AddPoints(a Point, b Point) Point {
    a.X += b.X
    a.Y += b.Y
    return a
}

func UnsetPoint() Point {
    return Point{-1,-1}
}
