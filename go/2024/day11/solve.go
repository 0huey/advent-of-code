package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	stones := Parse(os.Args[1])

	for range 25 {
		stones = Blink(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	fmt.Println("Part1:", sum)



	for range 50 {
		stones = Blink(stones)
	}

	sum = 0
	for _, count := range stones {
		sum += count
	}

	fmt.Println("Part2:", sum)
}

func Blink(old map[int]int) map[int]int {
	new := make(map[int]int)

	for stone, count := range old {
		str := strconv.Itoa(stone)

		if stone == 0 {
			new[1] += count

		} else if len(str) & 1 == 0 {
			left  := str[ : len(str)/2]
			right := str[len(str)/2 : ]

			leftn,  _ := strconv.Atoi(left)
			rightn, _ := strconv.Atoi(right)

			new[leftn] += count
			new[rightn] += count

		} else {
			new[stone*2024] += count
		}
	}
	return new
}

func Parse(filename string) map[int]int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	output := make(map[int]int)

	text := strings.TrimSpace(string(data))

	for _, num := range strings.Split(text, " ") {
		x, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		output[x]++
	}
	return output
}
