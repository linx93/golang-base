package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func showMessage(message string) {
	// -1
	defer waitGroup.Add(-1)
	fmt.Printf("exec showMessage i = %v\n", message)

}

/*func main() {
	fmt.Println(" main exec start")
	for i := 0; i < 10; i++ {
		go showMessage(strconv.Itoa(i))
		// +1
		waitGroup.Add(1)
	}

	//wait
	waitGroup.Wait()

	fmt.Println(" main exec end")
}
*/
