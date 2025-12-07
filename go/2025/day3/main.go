package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	banks := Parse(os.Args[1])

	Part1(banks)
	Part2(banks)
}

func Parse(filename string) [][]int {
	var banks [][]int

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))

	for _, line := range strings.Split(text, "\n") {
		var bank []int
		line = strings.TrimSpace(line)

		for _, battery := range line {
			b, err := strconv.Atoi(string(battery))
			if err != nil {
				panic(err)
			}
			bank = append(bank, b)
		}

		banks = append(banks, bank)
	}

	return banks
}

func Part1(banks [][]int) {
	joltage := 0

	for _, bank := range banks {
		num1 := 0
		num2 := 0

		for i, battery := range bank {
			if battery > num1 && i != len(bank)-1 {
				num1 = battery
				num2 = 0
			} else if battery > num2 {
				num2 = battery
			}
		}
		joltage += (num1 * 10) + num2
	}

	fmt.Println("Part1:", joltage)
}

const MAX_ACTIVE = 12

func Part2(banks [][]int) {
	joltage := 0

	for _, bank := range banks {
		var active []int

		for i, battery := range bank {
			if len(active) == 0 {
				active = append(active, battery)
				continue
			}

			appended := false
			for k := range active {
				if battery > active[k] && len(bank)-i >= MAX_ACTIVE-k {
					active = active[:k]
					active = append(active, battery)
					appended = true
					break
				}
			}
			if !appended && len(active) < MAX_ACTIVE {
				active = append(active, battery)
			}
		}

		bank_joltage := 0
		position := 1

		for _, battery := range slices.Backward(active) {
			bank_joltage += battery * position
			position *= 10
		}
		joltage += bank_joltage
	}

	fmt.Println("Part2:", joltage)
}
