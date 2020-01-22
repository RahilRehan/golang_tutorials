// concurrency is a design pattern
// parallelism is code running on multiple cores

//basic program to depict concurrency

package main

import(
	"fmt"
	"runtime"   //godoc.org/runtime
)

func main(){
	fmt.Println("Computer information: ")
	fmt.Println("Operating System: ", runtime.GOOS)
	fmt.Println("Architecture: ", runtime.GOARCH)
	cpuCount := runtime.NumCPU()
	fmt.Println("Number of CPU's available are: ", cpuCount)
	if(cpuCount > 1){
		fmt.Println("Parallelism possible on this computer, write concurrent code")
	}else{
		fmt.Println("Parallelism is not possible on this computer.")
	}

	//Goroutines
	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())

	go foo()
	bar()


	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())

	/* Note
		Only , BAR prints, because it takes some time for goroutine of foo to launch in time interval,
		our bar got executed and func main() completed, once func main is done, everything closes.
	*/

}

func foo(){
	for i:=0; i<10; i++{
		fmt.Println("FOO, ", i)
	}
}

func bar(){
	for i:=0; i<10; i++{
		fmt.Println("BAR, ", i)
	}
}
