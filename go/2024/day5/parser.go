package main

import (
	"os"
	"strings"
	"strconv"
)

type pageorder struct {
	first int
	second int
}

func Parse(filename string) (order []pageorder, updates [][]int) {
	data, err := os.ReadFile(filename)
	ErrorExit(err)

	str_data := strings.TrimSpace(string(data))

	sections := strings.Split(str_data, "\n\n")

	if len(sections) != 2 {
		panic("malformed challenge file")
	}

	order_lines := strings.Split(sections[0], "\n")

	for _, line := range order_lines {
		pages := strings.Split(line, "|")

		a, err := strconv.Atoi(pages[0])
		ErrorExit(err)

		b, err := strconv.Atoi(pages[1])
		ErrorExit(err)

		order = append(order, pageorder{first: a, second: b})
	}

	update_lines := strings.Split(sections[1], "\n")

	for _, line := range update_lines {
		var nums []int

		str_nums := strings.Split(line, ",")

		for _, x := range str_nums {
			num, err := strconv.Atoi(x)
			ErrorExit(err)

			nums = append(nums, num)
		}

		updates = append(updates, nums)
	}

	return
}

func ErrorExit(e error) {
	if e != nil {
		panic(e)
	}
}
