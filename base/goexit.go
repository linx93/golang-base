package main

import (
	"fmt"
	"runtime"
)

func show() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			//exit goroutine
			runtime.Goexit()
		}
	}
}

/*func main() {
	go show()
	time.Sleep(time.Second)
	println(runtime.NumCPU())
}*/

//print result
//i: 1
//i: 2
//i: 3
//i: 4
//i: 5
