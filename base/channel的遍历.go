package main

//定义int的channel
var chan1 chan int = make(chan int)

/*func main() {
	//开启一个goroutine
	go func() {
		for i := 0; i < 4; i++ {
			//写入channel
			chan1 <- i
		}
		//写完要关闭channel，否这读取时候，把channel中数据读完继续读就会报错【死锁】 fatal error: all goroutines are asleep - deadlock!
		close(chan1)
	}()
	for i := range chan1 {
		fmt.Println(i)
	}
}*/
