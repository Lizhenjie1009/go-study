package main

import "fmt"

/*
	管道	channel
		channel是一种类型，一种引用类型。	var 变量名 chan 元素类型
		管道是提供goroutine(协程)间的通讯方式，channel可以让一个goroutine发送特定的值到另一个goroutine的通信机制。

	golang的并发模型是CSP，提倡***通过通信共享内存***而不是***通过共享内存而实现通信***

	channel是一种特殊的类型，先入先出。声明channel的时候需要为其指定元素类型。
*/

func main() {
	var a chan int                    // 声明一个传递整型的管道
	var b chan bool                   // 声明一个传递布尔型的管道
	var c chan []int                  // 声明一个传递int切片的管道
	fmt.Printf("%T %T %T\n", a, b, c) // chan int chan bool chan []int

	// 1创建channel 容量3
	ch := make(chan int, 3)
	// 2管道存储值
	ch <- 10
	ch <- 20
	ch <- 30
	// 3获取管道的值
	a1 := <-ch
	fmt.Println(a1, <-ch) // 10 20
	a2 := <-ch
	fmt.Println(a2) // 30
	// 此时管道又空了
	ch <- 40
	ch <- 50
	// 4打印管道的长度和容量
	fmt.Printf("值:%v, 长度:%v, 容量:%v\n", ch, len(ch), cap(ch)) // 值:0xc00010c080, 长度:2, 容量:3
	// 5管道的类型 -- 引用数据类型
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64
	ch2 := ch1
	ch2 <- 44
	<-ch1
	<-ch1
	<-ch1
	b1 := <-ch1
	fmt.Println(b1) // 44
	// 6管道阻塞
	ch3 := make(chan int, 1)
	ch3 <- 10
	// ch3 <- 20 // 管道阻塞 deadlock! 死锁
	ch4 := make(chan string, 2)
	ch4 <- "a"
	ch4 <- "b"
	c1 := <-ch4
	c2 := <-ch4
	// c3 := <-ch4 // 管道阻塞 deadlock! 死锁
	fmt.Println(c1, c2)
}
