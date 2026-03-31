package main

import (
	"fmt"
)

func main() {
	for i := 65; i < 90; i++ {
		if i%2 == 0 {
			fmt.Printf("%c ", 'A'+i)
		} else {
			fmt.Printf("%c ", 'A'+i)
		}
	}
	fmt.Println()
}
