package main

import (
	"fmt"
	"io/ioutil"
	"os"

)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("usage : go run main.go <inputfine.txt> <outputfile.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	 inp, err := ioutil.ReadFile(inputFile)
	 if err != nil {
		fmt.Printf("Error reading this file please input the correct file: %v\n", err)
		return
	 }

	 err = ioutil.WriteFile(outputFile,inp, 0644)
	 if err != nil {
		fmt.Printf("Error finding this file : %v\n", err)
		return
	 }

	 fmt.Println("your program ran succsfuly ")
}