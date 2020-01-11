package main

import "fmt"

func main(){

	//slice initialization
	var x []int = []int{3,7,2,5,7}

	fmt.Println(x)

	fmt.Println("Length of array is ", len(x))
	fmt.Printf("Type of arrayay is %T\n", x)

	//sliced indexing
	fmt.Println("Sliced index is ", x[1:3])

	//appending elements
	y := []int{11,12,13}
	x = append(x, y...)  //... represents all the elements from y, unfirls all the elements from slice y to ints
	fmt.Println(x)

	//removing elements from slice
	x = append(x[:2], x[4:]...) // removing 2,3 from slice
	fmt.Println(x)


	//make(slize_type, no_of_initialized, capacity size)
	z := make([]int, 10, 100)
	fmt.Println(z)
	fmt.Println("Initialized length of slice is ",len(z)) //outputs only number of initialized locations
	fmt.Println("Total Capacity of array is ",cap(z))

	//note
	//z[11] = 345 //this wont work
	z = append(z, 345)
	fmt.Println(z, "\nLength is ", len(z))

	/*
	if you make an array with func make then you declare a capacity, after reaching which the array is doubled
	in case of normally declared array, array is copied everytime over with one element more size.
	*/
}
