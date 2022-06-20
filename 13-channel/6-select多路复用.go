package main

import (
	"fmt"
	"time"
)

/*
	select多路复用: 从多个管道获取数据，也可以用协程
*/

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hi " + fmt.Sprintf("%d", i)
	}

	// 并行读取数据
	// 1.开启协程并行
	// 2.select多路复用 -- 获取数据不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Println(v, "intChan")
			time.Sleep(time.Millisecond * 200)
		case v := <-stringChan:
			fmt.Println(v, "stringChan")
		default:
			fmt.Println("数据获取完毕")
			return
		}
	}
}
