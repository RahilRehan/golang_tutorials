/*
	Creating methods in go
*/
package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

//note
//func declaration syntax -> func (reciever) identifier(parameters) (return vals){}

//reciever is the struct to which we want to assign the method to
func (p person) speak(){ //any value of type p can now have speak method
	fmt.Println(p.first, p.last, p.age)
}

func main(){
	p1 := person{
		"rahil",
		"midde",
		21,
	}
	p1.speak()
}
