package main

import (
	"fmt"
)

func main() {

	// 	 Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.

	// These combinations are separated by a comma and a space.

	// Expected function
	// func PrintComb2() {

	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 98; j++ {
			fmt.Printf("%02d %02d", i, j)
			if i == 99 && j == 98 {
			} else {
				fmt.Print(",")
				fmt.Print(" ")
			}

		}
	}
	fmt.Println()
}
