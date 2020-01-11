/*
	Example to create multi dimensional arrays
*/

package main

import "fmt"

func main(){
	x := []string{"rah","hil","eah"}
	y := []string{"sha","iah","kah"}

	xy := [][]string{x,y}
	fmt.Println(xy)
	fmt.Println("Size of multi xy is ", len(xy))
}
