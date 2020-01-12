/*
	Concept of callback -> functional programming
	* run a function after a function is done running.
	* As a function is passed in as parameter to other function, later function runs only when the first function runs.
*/

package main
import "fmt"


//recieves as function as parameter
func even_sum(f func(ii ...int) []int, myL ...int) int{
	eL := f(myL...)
	sum := 0
	for _,v := range eL{
		sum += v
	}
	return sum
}


func main(){
	myL := []int{1,2,3,4,5,6,7,8,9}

	//creating a function expression to pass as an arguments to other function
	is_even := func(ii ...int) []int{
					evenL := []int{}
						for _,v := range ii{
							if(v%2==0){
								evenL = append(evenL, v)
							}
						}
						return evenL
				}

	fmt.Println(even_sum(is_even, myL...))
}
