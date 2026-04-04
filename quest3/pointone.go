package main

import (
	"fmt"
)

// Write a function that takes a pointer to an int as argument and gives to this int the value of 1.

// Expected function
func PointOne(n *int) {
	*n = 1

}

func main() {
	n := 0
	fmt.Print(&n)
	fmt.Println(n)
}