package main

import (
	"os"
	"strings"
	"strconv"
	"slices"
	"fmt"
)

func error_exit(e error) {
	if e != nil {
		panic(e)
	}
}

func Parse(filename string) (left, right []int) {
	data, err := os.ReadFile(filename)
	error_exit(err)

	lines := strings.Split(string(data), "\n")

	for i := range len(lines) {
		if len(lines[i]) == 0 {
			continue
		}

		nums := strings.Fields(lines[i])

		if len(nums) < 2 {
			fmt.Println("parser: bad line", i+1)
			continue
		}

		num, err := strconv.Atoi(nums[0])
		error_exit(err)

		left = append(left, num)

		num, err = strconv.Atoi(nums[1])
		error_exit(err)

		right = append(right, num)
	}

	slices.Sort(left)
	slices.Sort(right)

	return
}
