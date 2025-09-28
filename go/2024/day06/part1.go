package main

import (
    "os"
    "slices"
    "fmt"
)
func main() {
    argc := len(os.Args)

    if argc < 2 {
        fmt.Println("no file path specified")
        return
    }

    area, guard := Parse(os.Args[1])

    fmt.Println(guard)

    for i := range area {
        fmt.Println(area[i])
    }

    var visited []Point
    dir := Point{Y: -1, X: 0}

    for PointInArea(area, guard) {
        if !slices.Contains(visited, guard) {
            visited = append(visited, guard)
            area[guard.Y][guard.X] = AreaVisited
        }

        next := AddPoints(guard, dir)

        if !PointInArea(area, next) {
            break
        }

        if ObjectAtPoint(area, next) == AreaObstruction {
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
            guard = AddPoints(guard, dir)
        } else {
            guard = next
        }
    }
    fmt.Println("")
    for i := range area {
        fmt.Println(area[i])
    }
    fmt.Println(len(visited))
}
