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

	var sum_dist int = 0

	for i := range len(left) {
		sum_dist += abs(left[i] - right[i])
	}

	fmt.Println(sum_dist)
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
