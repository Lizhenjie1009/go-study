package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	读写互斥锁
		真正的互斥应该是读取和修改、修改和修改之间。读和读是没有互斥操作的必要。因此衍生出读写锁。
		读写锁可以让多个读操作并发，同时读取，但是对于写操作是完全互斥的。
		当一个goroutine(协程)进行写操作的时候，其他goroutine(协程)既不能读也不能写。
*/
var wg sync.WaitGroup

// var mutex sync.Mutex // 互斥锁
var mutex sync.RWMutex // 读写锁

func write() { // 只有一个协程可以写
	mutex.Lock()
	fmt.Println("执行写的操作")
	time.Sleep(time.Second)
	mutex.Unlock()
	wg.Done()
}
func read() { // 多个协程都可以读
	mutex.RLock()
	// mutex.Lock()
	fmt.Println("---执行读的操作")
	time.Sleep(time.Second)
	mutex.RUnlock()
	// mutex.Unlock()
	wg.Done()
}

func main() {
	// 并行加并发的任务

	// 开启10个协程执行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 开启10个协程执行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
}
