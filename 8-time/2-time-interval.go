package main

import (
	"fmt"
	"time"
)

// 时间间隔 - time-interval
func main() {
	/*
		time包中定义时间间隔类型的常量
			const (
				Nanosecond 		Duration = 1
				Microsecond						 = 1000 * Nanosecond
				Millisecond						 = 1000 * Microsecond
				Second								 = 1000 * Millisecond
				Minute								 = 60 * Second
				Hour									 = 60 * Minute
			)
	*/

	fmt.Println(time.Nanosecond)  // 1ns
	fmt.Println(time.Microsecond) // 1微妙
	fmt.Println(time.Millisecond) // 1毫秒
	fmt.Println(time.Second)      // 1秒
	fmt.Println(time.Minute)      // 1分钟
	fmt.Println(time.Hour)        // 1小时

	// 时间操作函数
	t := time.Now()
	fmt.Println(t)
	t = t.Add(time.Hour) // 增加一小时
	fmt.Println(t)

	// 定时器
	ticker := time.NewTicker(time.Second) // 1s的定时器
	n := 5
	for t := range ticker.C { // ticker.C 返回的只有一个参数 time
		n--
		fmt.Println(t)
		if n == 0 {
			ticker.Stop() // 停止定时器
			break
		}
	}

	fmt.Println(1)
	time.Sleep(time.Second) // 休眠一秒
	fmt.Println(1)
	time.Sleep(time.Second * 5) // 休眠五秒
	fmt.Println(1)

	// 每过一秒钟执行一次的死循环
	for {
		time.Sleep(time.Second)
		fmt.Println("定时任务")
	}
}
