package main

import (
	"fmt"
)

func main() {
	// 	Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

	// These combinations are separated by a comma and a space.
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for l := j + 1; l <= '9'; l++ {
				fmt.Printf("%c%c%c", i, j, l)
				if i == '7' && j == '8' && l == '9' {

				} else {
					fmt.Print(",")
					fmt.Print(" ")
				}
			}
		}
	}

	fmt.Println()

}
