package main

import (
	"fmt"
)

func PrintNbr(n int) {
	if n == 0 {
		fmt.Print("0")
		return
	}

	if n < 0 {
		fmt.Print("-")

		if n == -9223372036854775808 || n == -2147483648 {
			PrintNbr(-(n / 10))
			fmt.Print(rune(-(n % 10) + 0))
			return
		}
		n = -n
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	fmt.Print(rune(n%10 + 0))
}
func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	fmt.Println()
}

// my expected out sholud be
// -1230123
