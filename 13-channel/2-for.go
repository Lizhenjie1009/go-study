package main

import "fmt"

/*
	管道的循环 for range 没有key

	管道的关闭
*/

func main() {
	// rang1 	for range
	a := make(chan int, 3)
	a <- 1
	a <- 2
	a <- 3
	close(a)           // 循环遍历之前管道必须要关闭，否则deadlock!
	for v := range a { // 管道循环没有key
		fmt.Println(v)
	}

	// rang2	for
	a1 := make(chan string, 3)
	a1 <- "aaa"
	a1 <- "bbb"
	a1 <- "ccc"
	// 通过for循环变量channel的时候，管道可以不关闭, 建议关闭
	for i := 0; i < 3; i++ {
		fmt.Println(<-a1)
	}
}
