/*
	Implementing maps.
	unordered list for fast lookups
	similar to hash tables
*/

package main

import "fmt"

func main(){
	m := map[string]int{
		"rahil" : 21,
		"swaroop" : 20,
	}

	fmt.Println(m)
	fmt.Println("Age of Rahil is ", m["rahil"])

	//v is value returned from search and ok is bool depicting if the value is present or not.
	if v, ok := m["shaik"]; ok{
		fmt.Println("Found User with age", v)
	}else{
		fmt.Println("No such user's age found!")
	}

	//adding to maps and ranging over maps
	m["shaik"] = 18

	fmt.Println("All the elements in map are: ")
	for k, v := range m{
		fmt.Println(k, v)
	}

	//lets delete
	m["James"] = 32
	fmt.Println("Before deleting: ", m)
	delete(m, "James")
	fmt.Println("After deleting: ",m)
}
