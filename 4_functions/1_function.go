/*
	Function declarations, parameters, returns, multiple parameters
	passing arguments
*/

package main

import "fmt"

func myfunc(s string, b bool) (string, bool){
	return fmt.Sprint("Hello ", s, " your logic is "), b
}

func arfunc(x ...int){ //recieving number of ints as parameters
	fmt.Println(x)
	fmt.Printf("Type is %T\n", x)
}

func unfurling(x...int){
	fmt.Println(x)
	fmt.Printf("Type is %T\n", x)
}

func araspar(arr []int){
	fmt.Println("Recieved arr is ", arr)
	fmt.Printf("Type is %T", arr)
}

//main function
func main(){
	msg,boo := myfunc("rahil", true)
	fmt.Println(msg)
	fmt.Println(boo)
	arfunc(1,2,3,4)

	arr2 := []int{1,2,3,4,5,6}
	unfurling(arr2...) //unfirls integers from slice

	//sending and recieving array as argument, parameter.
	araspar(arr2)  //no unfirling or furling is required

}
