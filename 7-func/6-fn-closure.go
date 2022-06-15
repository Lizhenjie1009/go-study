package main

import "fmt"

// 闭包
func fn1() func() int {
	i := 10
	return func() int {
		return i + 1
	}
}
func fn2() func() int {
	i := 10
	return func() int {
		i = i + 1
		return i
	}
}

// 闭包-closure
func main() {
	a := fn1()
	fmt.Println(a()) // 11 -> return i + 1 没有更改i的值
	fmt.Println(a()) // 11 -> return i + 1 没有更改i的值
	fmt.Println(a()) // 11 -> return i + 1 没有更改i的值

	b := fn2()
	fmt.Println(b()) // 11 -> return i 返回修改i之后的值
	fmt.Println(b()) // 12 -> return i 返回修改i之后的值
	fmt.Println(b()) // 13 -> return i 返回修改i之后的值

}
