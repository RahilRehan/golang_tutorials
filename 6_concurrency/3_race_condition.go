//simulating race condition

//Race Condition: Diffreent GORoutines accessing same shared variables

package main

import(
	"fmt"
	"sync"
	"runtime"
)

func main(){

	const gs = 100
	var counter int = 0

	for i:=0; i<gs; i++{
		func(){
			v:= counter
			v+=1
			counter = v
		}()
	}

	fmt.Println("Without using goroutines: ", counter)

	counter = 0 //counter variable scope is in the main func, shared variable

	var wg sync.WaitGroup

	wg.Add(gs)

	for i:=0; i<gs; i++{
		go func(){
			v:= counter  //v variable scope is in the block of this function
			//time.Sleep(time.Second)
			runtime.Gosched() //race condition can be seen very well if you use this else it is hard to find.
			v+=1
			counter = v
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
