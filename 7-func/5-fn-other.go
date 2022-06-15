package main

import "fmt"

func fn(num int) {
	if num > 0 {
		fmt.Println(num)
		num--
		fn(num)
	}
}

// 递归1-100的和
func fn1(n int) int {
	if n > 1 {
		return n + fn1(n-1)
	}
	return 1
}

// 递归实现5的阶乘 1*2*3*4*5
func fn2(n int) int {
	if n > 1 {
		return n * fn2(n-1)
	}
	return 1
}

/*
	函数递归调用
*/
func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum) // 5050

	// 传入一个整数，递归1-整数之内的所有整数
	fn(2)

	// 递归1-100的和
	fmt.Println(fn1(100))

	// 递归实现5的阶乘 1*2*3*4*5
	fmt.Println(fn2(5))
}
