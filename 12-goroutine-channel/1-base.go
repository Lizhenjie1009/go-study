package main

import (
	"fmt"
	"sync"
	"time"
)

/*

	进程： 1个在执行的应用就是1个进程		(chrome)
	线程： 1个进程包含多个线程	(chrome-tab) 是进程的一个执行实例，是程序执行的最小单元，是比进程更小能独立运行的基本单位


	并发和并行是针对多线程程序而言的
		并发： 多个线程同时竞争一个位置，竞争到的才可以执行，每一个时间段只有一个线程在执行。
		并行： 多个线程可以同时执行，每一个时间段可以有多个线程同时进行。

		多线程程序在单核CPU上运行就是并发
		多线程程序在多核CPU上运行就是并行
		如果线程数大于CPU核数，即有并发又有并行
*/

/*
	并发-常规、例行 goroutine - channel
		golang中可以在主线程上起多个协程，多个协程可以实现并行或并发。

		协程：可以理解为用户级线程，类似线程和真正的线程不一样，这是对内核透明的，也就是系统不知道协程的存在，完全由用户自己程序进行调度的。
			golang的一大特色就是从语言层面原生支持协程，在函数或者方法前加go关键字就可以创建一个协程。


		多协程和多线程：golang中每个goroutine(协程)默认占用内存比Java、C的线程少。
			OS线程(操作系统线程)一般都有固定的栈内存(通常2MB左右)。
			一个goroutine(协程)占用内存非常小，只有2KB左右。

			多协程goroutine切换调度开销方面远比线程要少。

*/

// 主线程和goroutine同时运行
var wait sync.WaitGroup

func test() { // 协程50毫秒执行一次，共10次
	for i := 0; i < 10; i++ {
		fmt.Println("协程 test golang", i)
		time.Sleep(time.Millisecond * 50)
	}
	wait.Done() // 协程执行完，协程计数器减1
}

func main() {
	// 默认-主线程执行完就会结束，不管协程有没有执行完

	wait.Add(1) // 开启一个协程，协程计数器加1
	go test()   // 开启一个协程

	// 主线程50毫秒执行一次，共10次
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("golang", i)
	// 	time.Sleep(time.Millisecond * 50)
	// }
	// time.Sleep(time.Second) // 等待1秒中，让协程执行完 bad

	wait.Wait() // 等待所有协程执行完成, 结束主线程
	fmt.Println("主线程退出")
}
