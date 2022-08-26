package main

import (
	"fmt"
)

func showMsg(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Println(msg)
	}
}

/*func main() {
	go showMsg("java")
	for i := 0; i < 2; i++ {
		runtime.Gosched() //让出CPU时间片，重新等待安排任务
		fmt.Println("golang")
	}
	println("main end ")
}*/

//print result
//java
//java
//golang
//golang
// main end
