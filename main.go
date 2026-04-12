package main 

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("usage : go run main.go <inputfile.txt> < outputfile.txt>")
		return
	}
     // thia where i have decelerded the input 
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.ReadFile(input)
	if err !=  nil {
		fmt.Printf("error reading file please input the correct file%v\n", err)
		return
	}

	content := string(input)

	words := strconv.Fields(content)

	var result []string

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && 
	}



}