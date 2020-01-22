/*
	Program to expalain how to use mutexes to avoid race condition
	when, mutex is set(Locked), other goroutines cannot modify them.
*/


package main

import(
	"fmt"
	"sync"
	"runtime"
)

func main(){

	const gs = 100
	var counter int = 0 //counter variable scope is in the main func, shared variable

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)

	for i:=0; i<gs; i++{
		go func(){
			mu.Lock()
			v:= counter  //v variable scope is in the block of this function
			runtime.Gosched() //race condition can be seen very well if you use this else it is hard to find.
			v+=1
			counter = v
			mu.Unlock()
			wg.Done()
			fmt.Println("Number of goroutines: ", runtime.NumGoroutine())
		}()
	}

	wg.Wait()
	fmt.Println("With goroutines: ", counter)

	/*
		*Number of goroutines increase and decrease while printing, this is because of execution starts and comletes at different times.

		*Some goroutines mess up and due to race condition all of the goroutines will not get a chance to execute,
		hence counter is less than 100

		*Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.

		*to check if there is a race condition in the program, in cmd use the command `go run -race 3_race_condition.go`
	*/
}
