package main

import (
	"fmt"
)

// Write a function that prints all possible combinations of n different digits in ascending order.

// n will be defined as : 0 < n < 10

// Below are the references for the printing format expected.

// (for n = 1) '0, 1, 2, 3, ..., 8, 9'

// (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

func PrintCombN(n int) {
	var comb []int

	var generate func(start int)
	generate = func(start int) {
		if len(comb) == n {
			for _, v := range comb {
				fmt.Print(v)
			}
			fmt.Print(", ")
			return
		}

		for i := start; i < 10; i++ {
			comb = append(comb, i)
			generate(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	generate(0)
	fmt.Println()
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
