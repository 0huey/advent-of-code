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

	var order []pageorder
	var updates [][]int

	order, updates = Parse(os.Args[1])

	fmt.Println("page order:", order)
	fmt.Println("updates:", updates)

	var good_update bool
	var good_sum int = 0

	for _, update := range updates {
		good_update = true

		for i_num, num := range update {
			for _, entry := range order {

				if entry.first == num {
					// search update line backwards to make sure num isnt coming after entry.second

					for i_search := i_num-1; i_search >= 0; i_search-- {
						if update[i_search] == entry.second {
							good_update = false
							break
						}
					}
				}

				if !good_update {
					break
				}

				if entry.second == num {
					for i_search := i_num+1; i_search < len(update); i_search++ {
						if update[i_search] == entry.first {
							good_update = false
							break
						}
					}
				}
				if !good_update {
					break
				}
			}
			if !good_update {
				break
			}
		}

		if good_update {
			middle_value := update[ len(update) / 2]

			fmt.Println("good update", middle_value, update)

			good_sum += middle_value
		}
	}
	fmt.Println("sum of middle value from in-order updates:", good_sum)
}
