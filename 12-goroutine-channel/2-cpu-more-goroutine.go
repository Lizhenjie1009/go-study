package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	设置cpu
	开启多个协程
*/

var wg sync.WaitGroup

func run(i int) {
	defer wg.Done()
	fmt.Println("Hellow Goroutine", i)
}

func main() {
	// 获取当前计算机的cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	// 可以设置使用多个cpu
	// runtime.GOMAXPROCS(cpuNum - 1)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go run(i) // 开启协程
	}

	wg.Wait() // 等待所有协程结束
}
