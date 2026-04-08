package main

import(
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Println("usage: go run file.go <inputfile.txt> <outputFile.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	victor := []byte("Hello, victor how are you doing today hope all is we")
	_= os.WriteFile(inputFile,victor,0644)

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error Reading File  input the file: %v \n",err)
		return

	}

	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		fmt.Printf("Error out file not found!: %v\n", err)
	}
	fmt.Println("bravo the program printed without error")
}