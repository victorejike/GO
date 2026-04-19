// package main 

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// 	"strconv"
// )

// func main(){
// 	if len(os.Args) != 3 {
// 		fmt.Println("usage : go run main.go <inputfile.txt> < outputfile.txt>")
// 		return
// 	}
//      // thia where i have decelerded the input 
// 	inputFile := os.Args[1]
// 	outputFile := os.Args[2]

// 	input, err := os.ReadFile(input)
// 	if err !=  nil {
// 		fmt.Printf("error reading file please input the correct file%v\n", err)
// 		return
// 	}

// 	content := string(input)

// 	words := strconv.Fields(content)

// 	var result []string

// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(hex)" && i < 0 {

// 			decimal, err := strconv.ParseInt(words[i-1], 16, 64)

// 			if err != nil {
// 				result[len(words)-1] strconv.FormatInt(decimal, 10)
// 			} 
// 		} else {
// 			result = append(result, words[i])
// 		}

// 	} else if words[i] == "(bin)" &&  i < 0 {
// 		decimal, err := strconv.ParseInt(words[i-1], 2, 64)
// 		if err != nil {
// 			result[len(words)-1] strconv.FormatInt(decimal,10)
// 		} else{
// 			result = append
// 		}

// 	}



// }

package main 

import (
	"fmt"
	"os"
 	"strings"
	// "strconv" 
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("usage : go run main.go <inputfile> <outputfile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("file was read succssfuly")
		return
	}

	text := string(data)
	text = capitalized(text)

	err = os.WriteFile(outputFile,[]byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}
	fmt.Println("succssfuly ran the program")
}

// func capitalized(text string)string{
// 	result := strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
// 	return result
// }

// func transforma(text string)string{

// }

