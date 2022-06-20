package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	goroutine和channel协同
		并行执行
			一个写
			一个读
*/

var wg sync.WaitGroup

// 写数据
func write(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println(i, "写")
		time.Sleep(time.Millisecond * 500) // 写的慢，读的慢。管道是安全的
	}
	close(ch)
}

// 读数据
func read(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v, "读")
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	var ch = make(chan int, 10)
	wg.Add(1)
	go write(ch)
	wg.Add(1)
	go read(ch)

	wg.Wait()
	fmt.Println("主线程退出... ")
}
