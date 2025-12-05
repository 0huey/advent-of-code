package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProductRange struct {
	Start int
	End   int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	ranges := Parse(os.Args[1])

	Part1(ranges)
	Part2(ranges)
}

func Parse(filename string) []ProductRange {
	var products []ProductRange

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))

	for _, id_range := range strings.Split(text, ",") {
		var ids ProductRange

		id_split := strings.Split(id_range, "-")

		ids.Start, err = strconv.Atoi(id_split[0])
		if err != nil {
			panic(err)
		}

		ids.End, err = strconv.Atoi(id_split[1])
		if err != nil {
			panic(err)
		}

		products = append(products, ids)
	}
	return products
}

func Part1(products []ProductRange) {
	sum := 0
	for _, prod := range products {
		for i := prod.Start; i <= prod.End; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 != 0 {
				continue
			}
			mid := len(str) / 2

			if str[:mid] == str[mid:] {
				sum += i
			}
		}
	}
	fmt.Println("Part1:", sum)
}

func Part2(products []ProductRange) {
	sum := 0
	for _, prod := range products {
		for i := prod.Start; i <= prod.End; i++ {
			str := strconv.Itoa(i)

			for chars := 1; chars <= len(str)/2; chars++ {
				substr := str[:chars]

				invalid_id := true

				for compare := chars; compare < len(str); compare += chars {
					if compare+chars > len(str) || substr != str[compare:compare+chars] {
						invalid_id = false
						break
					}
				}

				if invalid_id {
					sum += i
					break
				}
			}
		}
	}

	fmt.Println("Part2:", sum)
}
