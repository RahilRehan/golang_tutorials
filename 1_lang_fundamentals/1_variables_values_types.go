//line below tells that this file is a main package, and this must be executed first.
package main

//other packages that you want to import are specified below
import "fmt"

//main function, execution of program is started and ended here.
func main(){
	//clear syntax

	//syntax: keyword identifier type = value
	var x int = 34  //initialization
	fmt.Println(x)

	//shorthand declaration
	y := 43  //automatic type identification and assignment
	fmt.Println(y)

	//other types
	s := "rahil"
	fmt.Println(s)

	var b bool = true
	fmt.Println(b)

	hello := "hello, world!"
	fmt.Println(hello)

	//note

	//declaration : only decclaring type of var
	//assignemnt: storing of value to a variable
	//initialization: assignment of a value to a variable at the time of declaration

	//only declared values are assignes to 0 in case of int, false in case of bool and "" in case of string.
}
