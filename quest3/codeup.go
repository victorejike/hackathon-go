package main

import(
	"fmt"
	"strings"
)

func Capitalized(word []string)[]string{
	 var result []string
	 for i := 0; i < len(word); i++{
		if word[i] == "(UP)" {
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}

		} else {
			result = append(result, word[i])
		}
	 }
	 return result
}

func main(){
	fmt.Println(Capitalized([]string{"go", "is", "fun", "(UP)"}))
	fmt.Println(Capitalized([]string{"hello", "(UP)"}))
}