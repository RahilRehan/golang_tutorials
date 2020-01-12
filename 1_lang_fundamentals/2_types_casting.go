package main
import "fmt"

//creating a new type with underlying type
//keyword identifier underlying_type
type hotdog int

func main(){
	var a int = 23
	var z hotdog = 43

	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("Type of hotdog is %T\n", z)

	//type conversion
	//a = hotdog //not possible
	a = int(z)
	fmt.Println("Value of a is ", a)
	fmt.Printf("Type of a is %T\n", a)
}
