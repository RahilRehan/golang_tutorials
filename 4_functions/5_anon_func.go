/*
	Anonymous functions.
	In below example we ran an anonymous fucntion and returned it at the same time.
*/

package main

import "fmt"

func foo(x int) int{
	return func(x int) int{
		x+=5
		return x
	}(x)
}

func main(){
	x := 34
	fmt.Println(foo(x))
}
