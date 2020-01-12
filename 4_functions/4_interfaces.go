/*
	Interfaces: **yet to write**
	***any value can have more than one type***
*/

package main

import (
	"fmt"
	"math"
)

type circle struct{
	radius float64
}

type square struct{
	length float64
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}

func (s square) area() float64{
	return s.length*s.length
}

//any type that has method area() is also a type shape
//***any value can have more than one type***
type shape interface{
	area() float64
}

func info(s shape) float64{
	return s.area()
}

func main(){
	ci := circle{
		radius: 4.8,
	}

	sq := square{
		length: 8.2,
	}

	fmt.Println(info(ci))
	fmt.Println(info(sq))

}
