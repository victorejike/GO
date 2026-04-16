package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("usage : go run vic.go <inputfile> <outputfile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error cant fine file please input file %v\n", err)
		return
	}

	err = ioutil.WriteFile(outputFile,content,0644)
	if err != nil {
		fmt.Printf("Error writing this file please provide another file dear%v\n", err)
		return
	}

	fmt.Println("your program has run succssfully")
}