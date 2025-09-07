package main

import (
    "os"
    "slices"
    "fmt"
)

func ErrorExit(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    argc := len(os.Args)

    if argc < 2 {
        fmt.Println("no file path specified")
        return
    }

    area, guard, err := Parse(os.Args[1])
    ErrorExit(err)

    var visited []Point
    dir := Point{Y: -1, X: 0}

    for PointInArea(area, guard) {
        next := AddPoints(guard, dir)
        char, _ := CharAtPoint(area, next)

        if char == "#" {
            // rotate dir 90 deg
            if dir.Y == -1 {
                dir = Point{Y: 0, X: 1}
            } else if dir.Y == 1 {
                dir = Point{Y: 0, X: -1}
            } else if dir.X == -1 {
                dir = Point{Y: -1, X: 0}
            } else if dir.X == 1 {
                dir = Point{Y: 1, X: 0}
            }
        } else {
            if !slices.Contains(visited, guard) {
                visited = append(visited, guard)
            }
            guard = next
        }
    }
    fmt.Println(len(visited))
}
