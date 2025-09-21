package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operator int

const (
	Add Operator = iota
	Mul
	Concat
	NumOperators
)

type Chal struct {
	Answer int
	Values []int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no file specified")
		return
	}

	data, err := Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	var sum int
	for _, c := range data {
		ops := slices.Repeat([]Operator{Add}, len(c.Values)-1)
		if Recurse(c, ops) {
			sum += c.Answer
		}
	}
	fmt.Println(sum)
}

func Recurse(data Chal, ops []Operator) bool {
	sum := data.Values[0]
	var err error

	for i, o := range ops {
		switch o {
		case Add:
			sum += data.Values[i+1]
		case Mul:
			sum *= data.Values[i+1]
		case Concat:
			strnum := strconv.Itoa(sum) + strconv.Itoa(data.Values[i+1])
			sum, err = strconv.Atoi(strnum)
			if err != nil {
				panic(err)
			}
		}
	}

	if sum == data.Answer {
		return true
	}

	for i := range slices.Backward(ops) {
		ops[i]++
		if ops[i] == NumOperators && i == 0 {
			return false
		} else if ops[i] == NumOperators {
			ops[i] = 0
		} else {
			break
		}
	}
	return Recurse(data, ops)
}

func Parse(filename string) ([]Chal, error) {
	var parsed []Chal

	data, err := os.ReadFile(filename)
	if err != nil {
		return parsed, err
	}

	data2 := strings.TrimSpace(string(data))

	for line := range strings.Lines(data2) {
		var c Chal

		linesplit := strings.Split(line, ":")

		c.Answer, err = strconv.Atoi(linesplit[0])
		if err != nil {
			return parsed, err
		}

		vals := strings.Split(strings.TrimSpace(linesplit[1]), " ")

		for _, n := range vals {
			x, err := strconv.Atoi(n)
			if err != nil {
				return parsed, err
			}
			c.Values = append(c.Values, x)
		}
		parsed = append(parsed, c)
	}
	return parsed, nil
}
