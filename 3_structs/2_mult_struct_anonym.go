/*
	This example exaplains structs containing structures
	and also creation of anonymous structures
*/

package main

import "fmt"

type vehicle struct{
	doors int
	color string
}

//truck and sedan also has properties of vehicle

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

func main(){
	tr := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}
	sed := sedan{
		vehicle: vehicle{
			doors:4,
			color:"black",
		},
		luxury:false,
	}
	fmt.Println(tr)
	fmt.Println(sed)

	//anonymous structures
	p := struct{
		first string
		age int
	}{
		first: "Donny",
		age:14,
	}
	fmt.Println(p)
}
