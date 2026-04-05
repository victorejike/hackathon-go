package main 


import (
	"fmt"
)

func DivMod(a int, b int, div *int, mod *int) {
// 	This function will divide the int a and b.

// The result of this division will be stored in the int pointed by div.

// The remainder of this division will be stored in the int pointed by mod.

  if b == 0 {
	return
  }

  *div =  a/b
  *mod = a % b

   
}

func main() {
	a := 13
	b := 2
	var div int
	var mod int
     DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}