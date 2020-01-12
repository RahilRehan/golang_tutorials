/*
	Function Expressions: assigning functions to a variable and calling variable with arguments
*/

package main

import "fmt"

func main(){
	f := func(ii...int) int{
			pr := 1
			for _,v := range ii{
				pr*=v
			}
			return pr
		}

	iv := []int{1,2,3,4}
	fmt.Println(f(iv...))
}
