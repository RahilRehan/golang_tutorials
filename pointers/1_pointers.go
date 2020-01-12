/*
	Pointer examples,
	* -> dereference
	& -> get address
*/

package main
import "fmt"

type person struct{
	name string
}

// func (p *person) changeMe(){
// 	p.name = "rehan"
// }

//other way
func changeMe(p *person){
	p.name = "rehan"
}

func main(){
	// var x int = 45
	// fmt.Println("Value of x is ", x)
	// fmt.Println("Address of x is ", &x)

	p1 := person{
		name: "rahil",
	}

	//p1.changeMe()
	changeMe(&p1)
	fmt.Println(p1.name)

}
