/*
	Explanation of keyword defer in GOlang
*/

package main

import "fmt"

func foo(){
	fmt.Println("This is function foo!")
}

func bar(){
	fmt.Println("This is function bar!")
}

func main(){
	defer foo()  //runs at the end of function it is being called from(main, here)
	bar()
}
