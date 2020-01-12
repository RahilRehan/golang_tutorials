//structures in go

package main

import "fmt"


//this struct is a blue print like a class to store data
type person struct{
	first string
	last string
	favFlavours []string
}

func main(){
	//value of type(go's terminology) ->OOPs terminology: instance of class
	p1 := person{
		first: "shaik",
		last: "khan",
		favFlavours: []string{"lol", "col", "mcv"},
	}

	p2 := person{
		first: "rahil",
		last: "rehan",
		favFlavours: []string{"str", "chc", "jvn"},
	}

	//map to store all the values created from struct
	m := map[string]person{
		p1.last:p1,
		p2.last:p2,

	}

	fmt.Println(m)
	fmt.Println(p1)
	fmt.Println(p2)
}
