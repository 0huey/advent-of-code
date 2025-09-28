package main

import (
    "os"
    "slices"
    "fmt"
)

type Redirect struct {
    Pos Point
    Dir Point
}

func main() {
    argc := len(os.Args)

    if argc < 2 {
        fmt.Println("no file path specified")
        return
    }

    area, start := Parse(os.Args[1])

    var route []Point
    guard := start
    dir := Point{Y: -1, X: 0}

    for PointInArea(area, guard) {
        if !slices.Contains(route, guard) {
            route = append(route, guard)
        }
        next := AddPoints(guard, dir)
        if !PointInArea(area, next) {
            break
        }
        if ObjectAtPoint(area, next) == AreaObstruction {
            dir = RotateGuard(dir)
        } else {
            guard = next
        }
    }

    var loops int

    for _, insert := range route {
        if insert == start {
            continue
        }
        guard = start
        dir := Point{Y: -1, X: 0}
        var redirects []Redirect

        old_area_point := ObjectAtPoint(area, insert)
        SetObjectAtPoint(area, insert, AreaObstruction)

        for PointInArea(area, guard) {
            next := AddPoints(guard, dir)
            if !PointInArea(area, next) {
                break
            }
            if ObjectAtPoint(area, next) == AreaObstruction {
                redir := Redirect{Pos: guard, Dir: dir}
                if slices.Contains(redirects, redir) {
                    loops++
                    break;
                }
                redirects = append(redirects, redir)

                dir = RotateGuard(dir)
            } else {
                guard = next
            }
        }
        SetObjectAtPoint(area, insert, old_area_point)
    }
    fmt.Println(loops)
}

func RotateGuard(dir Point) Point {
    if dir.Y == -1 {
        dir = Point{Y: 0, X: 1}
    } else if dir.Y == 1 {
        dir = Point{Y: 0, X: -1}
    } else if dir.X == -1 {
        dir = Point{Y: -1, X: 0}
    } else if dir.X == 1 {
        dir = Point{Y: 1, X: 0}
    }
    return dir
}
