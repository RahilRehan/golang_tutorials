/*
	Example to show how to return function from function.
	hint: use function expressions
*/

package main
import "fmt"

func foo() func(x int) int{
	f := func(x int) int{
		return x+45
	}
	return f
}

func main(){
	f := foo()
	fmt.Println(f(34))
}
