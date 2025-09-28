package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"slices"
)

type FileType int
const (
	Type_File FileType = iota
	Type_Space
)

type FileBlock struct {
	Type FileType
	Id int
}

type FileSystem []FileBlock

type FileBlock2 struct {
	Type FileType
	Id int
	Size int
}

type FileSystem2 []FileBlock2

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing input file")
		os.Exit(1)
	}

	chal := Parse(os.Args[1])

	left  := 0
	right := len(chal) - 1;

	for left < right {
		if chal[left].Type != Type_Space {
			left++

		} else if chal[right].Type != Type_File {
			right--

		} else {
			temp := chal[left]
			chal[left] = chal[right]
			chal[right] = temp
			left++
			right--
		}
	}

	var checksum int

	for i, block := range chal {
		checksum += i * block.Id
	}

	fmt.Println("Part 1:", checksum)


	chal2 := Parse2(os.Args[1])

	right = len(chal2) - 1

	for right > 0 {
		right_obj := chal2[right]

		if right_obj.Type == Type_Space {
			right--

		} else if right_obj.Type == Type_File {

			inserted := false

			for left := 0; left < right; left++ {
				left_obj := chal2[left]

				if left_obj.Type != Type_Space {
					continue
				}

				if left_obj.Size == right_obj.Size {
					chal2[right] = left_obj
					chal2[left]  = right_obj
					break
				}

				if left_obj.Size > right_obj.Size {
					chal2[left].Size -= right_obj.Size
					chal2[right] = left_obj
					chal2[right].Size = right_obj.Size
					chal2 = slices.Insert(chal2, left, right_obj)
					inserted = true
					break
				}
			}
			if !inserted {
				right--
			}
		}
	}

	var pos int
	var checksum2 int

	for _, block := range chal2 {
		for range block.Size {
			if block.Type == Type_File {
				checksum2 += block.Id * pos
			}
			pos++
		}
	}

	fmt.Println("Part 2:", checksum2)
}

func Parse(filename string) FileSystem {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))

	var fs FileSystem
	var this_type FileType = Type_File
	var file_id int = 0

	for _, char := range text {
		size, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		block := FileBlock{
			Type: this_type,
		}

		if this_type == Type_File {
			block.Id = file_id

			file_id++
			this_type = Type_Space

		} else {
			this_type = Type_File
		}

		for range size {
			fs = append(fs, block)
		}
	}

	return fs
}

func Parse2(filename string) FileSystem2 {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := strings.TrimSpace(string(data))

	var fs FileSystem2
	var this_type FileType = Type_File
	var file_id int = 0

	for _, char := range text {
		size, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}

		block := FileBlock2{
			Type: this_type,
			Size: size,
		}

		if this_type == Type_File {
			block.Id = file_id

			file_id++
			this_type = Type_Space

		} else {
			this_type = Type_File
		}

		fs = append(fs, block)
	}

	return fs
}

// this will only work for the example where the file Ids <= 9
func (fs FileSystem) Print() {
	for _, block := range fs {

		if block.Type == Type_File {
			fmt.Print(strconv.Itoa(block.Id))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func (fs FileSystem2) Print() {
	for _, block := range fs {

		if block.Type == Type_File {
			fmt.Print(strings.Repeat(strconv.Itoa(block.Id), block.Size))
		} else {
			fmt.Print(strings.Repeat(".", block.Size))
		}
	}
	fmt.Println()
}

