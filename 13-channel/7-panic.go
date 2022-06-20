package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Hello, world")
	}
}
func test() {
	defer func() { // 捕获错误
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()

	var myMap map[int]string // 错误代码，没有分配储存空间
	myMap[0] = "golang"
}
func main() {
	go sayHello()
	go test()
	time.Sleep(time.Second)
}
