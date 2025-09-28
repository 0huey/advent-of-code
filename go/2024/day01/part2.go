package main

import (
	"fmt"
	"os"
)

func main() {
	argc := len(os.Args)

	if argc < 2 {
		fmt.Println("no file path specified")
		return
	}

	left, right := Parse(os.Args[1])

	right_counts := make(map[int]int)

	for _, num := range right {
		right_counts[num]++
	}

	var score int = 0

	for _, num := range left {
		score += num * right_counts[num]
	}

	fmt.Println(score)
}
