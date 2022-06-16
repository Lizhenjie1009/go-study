package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	统计1-10000000数字中的素数，并打印这些素数
	素数：除了1和本身不能被其他数整除的数
	素数：n以内的数整除 n%i==0 && n>i
*/

// 协程1 统计 1-30000
// 协程2 统计 30001-60000
// 协程3 统计 60001-90000
// 协程4 统计 90001-120000

var wg sync.WaitGroup

// start: (n-1)*30000+1 	end: n*30000
func test(n int) {
	defer wg.Done()
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num == 1 {
			break
		}
		flag := true
		for i := 2; i < num; i++ { // 除了num以内的数
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num, "是素数")
		}
	}
}

func main() {
	// start := time.Now().Unix()
	// // for循环计算12w素数
	// for num := 2; num < 120000; num++ {
	// 	flag := true
	// 	for i := 2; i < num; i++ { // num以内的数相除
	// 		if num%i == 0 { // num不是素数，下一个num
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		fmt.Println(num, "是素数")
	// 	}
	// }
	// end := time.Now().Unix()
	// fmt.Println(end - start) // 8毫秒

	start1 := time.Now().Unix()
	// 使用协程计算12w素数
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end1 := time.Now().Unix()
	fmt.Println(end1 - start1) // 4毫秒

}
