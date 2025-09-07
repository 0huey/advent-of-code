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

func Parse(filename string) ([]string, Point, error) {
    var lines []string
    var guard Point

    data, err := os.ReadFile(filename)
    if err != nil {
        return lines, guard, err
    }

    str_data := strings.TrimSpace(string(data))
    lines = strings.Split(str_data, "\n")

    line_len := len(lines[0])

    for lineno := range lines {
        if len(lines[lineno]) != line_len {
            return lines, guard, fmt.Errorf("Mismatched line lengths at line %d", lineno)
        }

        i := strings.Index(lines[lineno], "^")
        if i >= 0 {
            guard = Point{X: i, Y: lineno}
        }
    }
    return lines, guard, nil
}

func PointInArea(area []string, p Point) bool {
    return p.X >= 0 && p.Y >= 0 && p.Y < len(area) && p.X < len(area[0])
}

func CharAtPoint(area []string, p Point) (string, error) {
    if !PointInArea(area, p) {
        return "", fmt.Errorf("Point outside area")
    }
    return string(area[p.Y][p.X]), nil
}

func AddPoints(a Point, b Point) Point {
    a.X += b.X
    a.Y += b.Y
    return a
}
