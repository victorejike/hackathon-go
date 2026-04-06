package main

import (
	"fmt"
)

func PrintComb(n int) {
	var comb []int

	var gen func(start int)
	gen = func(start int) {
		if len(comb) == n {
			for _, v := range comb {
				fmt.Print(v)
			}
			fmt.Print(", ")
			return
		}

		for i := start; i < 10; i++ {
			comb = append(comb, i)
			gen(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	gen(0)
	fmt.Println()
}

func main() {
	PrintComb(1)
	PrintComb(3)
	PrintComb(9)

}
