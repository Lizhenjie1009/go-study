package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	编译：
		go build 文件名
		go build -race 文件名 // 查看竞争关系
*/

/*
	互斥锁
		互斥锁是传统并发编程中对共享资源进行访问控制的主要手段。
		它由标准库sync中的Mutex结构体类型表示。sync.Mutex类型只有两个公开的指针方法，Lock和Unlock。
			Lock: 锁定当前共享的资源
			Unlock： 进行解锁
*/

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex // 生成互斥锁
var m = make(map[int]int, 0)

func test() {
	defer wg.Done()
	mutex.Lock() // 访问同一资源的时候，加锁，只能一个协程操作
	count++      // 同一时刻可能有多个协程去操作count, 所以需要互斥锁
	fmt.Println("the count is ", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock() // 访问结束解锁
}

func opMap(num int) { // 求num的阶乘
	defer wg.Done()
	mutex.Lock()
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum // 可能有多个协程同时去操作m，产生竞争
	fmt.Println(m)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
}

func main() {
	for i := 0; i < 40; i++ {
		wg.Add(1)
		// go test()
		go opMap(i)
	}
	wg.Wait()
}
