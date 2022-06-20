package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 1.往intChan中放1-120000的数
func putNum(intChan chan int) {
	defer wg.Done()
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 2.从intChan取数据，判断是否是素数。并放在primeChan
func primeNum(intChan, primeChan chan int, exitChan chan bool) {
	defer wg.Done()
	for num := range intChan {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 { // num不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	// close(primeChan) // 开启了16个协程，给一个协程的channel关闭了，其他协程没办法写数据了
	exitChan <- true // 没执行完一次存放一条数据
}

// 4.printPrime打印素数
func printPrime(primeChan chan int) {
	defer wg.Done()
	for v := range primeChan {
		fmt.Println(v)
	}
}

func main() {
	// runtime.GOMAXPROCS(4)

	intChan := make(chan int, 1000)   // 存放数字的1-120000
	primeChan := make(chan int, 1000) // 存放素数
	exitChan := make(chan bool, 16)   // 标识primeChan close

	start := time.Now().Unix()

	wg.Add(1)
	go putNum(intChan)

	for i := 0; i < 16; i++ { // 开启16个协程统计素数
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 3.判断exitChan是否存满16个true
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 16; i++ {
			<-exitChan // 上面没执行完，会在这里排队。全部输出完，执行完关闭primeChan管道
		}
		close(primeChan)
	}()

	wg.Add(1)
	go printPrime(primeChan)

	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end-start, "s") // 2s
}
