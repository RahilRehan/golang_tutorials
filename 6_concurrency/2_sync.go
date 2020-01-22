/*
	This program is to make our main func wait till the goroutine foo is done executing,
	we use waitgroup type from sync package to achieve this.

	goroutines doc:
	https://golang.org/doc/effective_go.html#goroutines

*/

package main

import(
	"fmt"
	"runtime"   //godoc.org/runtime
	"sync"      //https://godoc.org/sync
)

var wg sync.WaitGroup  //package level scope

func main(){

	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())

	wg.Add(1) //add one thing to wait
	go foo()
	bar()


	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())

	wg.Wait() //wait for something to be done.
}

func foo(){
	for i:=0; i<10; i++{
		fmt.Println("FOO, ", i)
	}
	wg.Done() //tell that foo is done executing
}

func bar(){
	for i:=0; i<10; i++{
		fmt.Println("BAR, ", i)
	}
}

/*
 Note that we have multiple cores and also our code is written concurrently,
 but the goroutine of main func runs faster as it takes a bit of a time for goroutine foo starts.
 These goroutines run is parallel when there are lot of operations increasing.
*/
