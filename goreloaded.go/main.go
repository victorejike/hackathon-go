package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("usage : go run main.go <inputfile.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error cant fine this kind of file: %v\n", err)
		return
	}

	test := 

	err = ioutil.WriteFile(outputFile,[]byte(test), 0644)
	if err != nil {
		fmt.Printf("Error finding file: %v\n", err)
		return
	}
	fmt.Println("Bim")
}


func convertHexToBin(input string)string{
	words := strings.Fields(input)
	 for i := 0; i < len(input); i++ {
		if words[i] == "(hex)" && i > 0 {
			deciaml,err := strconv.ParseInt(input, 16, 64)
		}
	 }
}