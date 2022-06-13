package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

// 函数参数简写，相同类型可以简写
func sub(x, y int) int {
	return x - y
}

// 函数的可变参数 ...
func mulargs(x ...int) int { // 切片 x
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func mulargs1(x int, y ...int) int { // x=1  y=切片
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

// 返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名：函数定义时可以给返回值命名
// 在函数体内可以使用返回值命名变量
// 最后直接return
func calc1(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	a := sum(2, 3)
	fmt.Println(sum(1, 2), a)
	fmt.Println(mulargs(1, 2, 3, 4))
	fmt.Println(mulargs1(1, 2, 3, 4, 5))
	fmt.Println(calc(2, 3))
	x, y := calc1(5, 3)
	_, b := calc2(5, 3)
	fmt.Println(x, y, b)
}
