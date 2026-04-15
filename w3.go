// package main 

// import(
// 	"fmt"
// )

// func main(){

// 	var i, j = 19, "helo"
// 	fmt.Print(i ,"\n")
// 	fmt.Print(j, "\n")
// }

package main 

import (
	"fmt"
)

func main(){

	var i = 15.5
	var text = "hello world!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", text)
	fmt.Printf("%#v\n", text)
	fmt.Printf("%v%%\n", text)
}