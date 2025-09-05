package main

import (
	"fmt"
	"os"
	"math/rand"
)

func main() {
	argc := len(os.Args)

	if argc < 2 {
		fmt.Println("no file path specified")
		return
	}

	var order []pageorder
	var updates [][]int

	order, updates = Parse(os.Args[1])

	var bad_updates [][]int

	for _, u := range updates {
		if !IsInOrder(u, order) {
			bad_updates = append(bad_updates, u)
		}
	}

	var middle_sum int = 0

	middle_values := make(chan int, len(bad_updates))

	for _, bad := range bad_updates {
		go FixAndGetMiddle(bad, order, middle_values)
	}

	for elem := range middle_values {
		fmt.Println(elem)
		middle_sum += elem
	}

	fmt.Println(middle_sum)
}

func FixAndGetMiddle(bad []int, order []pageorder, middle chan<- int) {
	// bogo sort
	for !IsInOrder(bad, order) {
		for i := range bad {
			j := rand.Intn(i+1)
			bad[i], bad[j] = bad[j], bad[i]
		}
	}
	middle <- bad[len(bad) / 2]
}

func IsInOrder(update []int, order []pageorder) (bool) {
	for i_num, num := range update {
		for _, entry := range order {

			if entry.first == num {
				// search update line backwards to make sure num isnt coming after entry.second

				for i_search := i_num-1; i_search >= 0; i_search-- {
					if update[i_search] == entry.second {
						return false
					}
				}
			}

			if entry.second == num {
				for i_search := i_num+1; i_search < len(update); i_search++ {
					if update[i_search] == entry.first {
						return false
					}
				}
			}
		}
	}

	return true
}
