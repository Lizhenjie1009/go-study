package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	管道默认情况是双向的
	单向管道，只读、只写
		只读: var ch1 <-chan int
		只写: var ch2 chan<- int
*/

func main() {
	// 默认
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch, <-ch)

	// ch1 := make(<-chan int, 2) // 只读
	ch2 := make(chan<- int, 2) // 只写
	ch2 <- 11
	ch2 <- 22

	wg.Add(1)
	go write(ch)
	wg.Add(1)
	go read(ch)

	wg.Wait()
	fmt.Println("线程执行完毕!")
}

var wg sync.WaitGroup

func write(ch chan<- int) { // 只写的管道
	wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println(i, "写")
		time.Sleep(time.Millisecond * 200)
	}

	close(ch) // 执行完管道操作，关闭管道
}
func read(ch <-chan int) { // 只读的管道
	defer wg.Done()
	for v := range ch {
		fmt.Println(v, "读")
	}
}
