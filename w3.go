// package main 

// import(
// 	"fmt"
// )

// func main(){

// 	var i, j = 19, "helo"
// 	fmt.Print(i ,"\n")
// 	fmt.Print(j, "\n")
// }

// 

package main 

import (
	"fmt"
)

func main(){
	var a = 15
 // this for formatspecifer to to print in go
	fmt.Printf("% X\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%+d\n", a)
	fmt.Printf("%o\n", a)
	fmt.Printf("%O\n", a)
}