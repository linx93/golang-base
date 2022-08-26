package main

import (
	"fmt"
	"sync"
)

//加锁demo

var globalValue = 100

//go的互斥锁
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	// wg.Add(-1)等价于wg.Done()
	//defer wg.Add(-1)
	defer wg.Done()
	lock.Lock()
	globalValue += 1
	fmt.Printf("globalValue++ = %v\n", globalValue)
	lock.Unlock()
}

func sub() {
	defer wg.Done()
	lock.Lock()
	globalValue -= 1
	fmt.Printf("globalValue-- = %v\n", globalValue)
	lock.Unlock()
}

//func main() {
//	for globalValue := 0; globalValue < 100; globalValue++ {
//		wg.Add(1)
//		go add()
//		wg.Add(1)
//		go sub()
//	}
//	wg.Wait()
//	fmt.Printf("main end  i=%v\n", globalValue)
//}
