package main

import "fmt"

func main(){
	fmt.Println("Welcome Arrays!")
	var arr [1]int  //declaration of array of size 1, with values 0
	fmt.Println(arr[0])

	//arrays are also called as slices in golang,
	//line below can be read as creating arr2 using a slice of ints
	// arr2 := []int{5,4,3,2,1}
	var arr2 []int
	arr2 =  []int{5,9,3,7,4}

	fmt.Println("Array is ", arr2)
	fmt.Println("Size of array is ", cap(arr2))

	//loop over elements
	// for i:=0; i<5; i++{
	// 	fmt.Println(arr2[i])
	// }

	//looping with index, similar to enumerate in python
	fmt.Println("Elements at indices are: ")
	for i, e := range arr2{
		fmt.Println(i, e)
	}

}
