package main

import (
	"fmt"
	"os"
	"strings"
)

type PaperMap []string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	paper := Parse(os.Args[1])

	Part1(paper)
	Part2(paper)
}

func Parse(filename string) PaperMap {
	var paper PaperMap

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))

	for i, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)

		if i > 0 && len(paper[0]) != len(line) {
			panic(fmt.Errorf("line len differs at line %d", i+1))
		}

		paper = append(paper, line)
	}
	return paper
}

func (p PaperMap) ValidIndex(row int, col int) bool {
	return row >= 0 && row < len(p) && col >= 0 && col < len(p[0])
}

func Part1(paper PaperMap) {
	var diff [3]int = [3]int{-1, 0, 1}

	accessable := 0

	for row := range paper {
		for col := range paper[row] {
			if paper[row][col] == '@' {
				neighbors := 0

				for _, x := range diff {
					for _, y := range diff {
						if x == 0 && y == 0 {
							continue
						}
						rowx, coly := row+x, col+y

						if paper.ValidIndex(rowx, coly) && paper[rowx][coly] == '@' {
							neighbors++
						}
					}
				}
				if neighbors < 4 {
					accessable++
				}
			}
		}
	}
	fmt.Println("Part1:", accessable)
}

func Part2(paper PaperMap) {
	var diff [3]int = [3]int{-1, 0, 1}

	removed_count := 0
	removed := true

	for removed {
		removed = false
		var new_paper PaperMap

		for row := range paper {
			var paper_row string

			for col := range paper[row] {
				if paper[row][col] == '.' {
					paper_row += "."

				} else if paper[row][col] == '@' {
					neighbors := 0

					for _, x := range diff {
						for _, y := range diff {
							if x == 0 && y == 0 {
								continue
							}
							rowx, coly := row+x, col+y

							if paper.ValidIndex(rowx, coly) && paper[rowx][coly] == '@' {
								neighbors++
							}
						}
					}
					if neighbors < 4 {
						removed_count++
						removed = true
						paper_row += "."
					} else {
						paper_row += "@"
					}
				}
			}
			new_paper = append(new_paper, paper_row)
		}
		paper = new_paper
	}
	fmt.Println("Part2:", removed_count)
}
