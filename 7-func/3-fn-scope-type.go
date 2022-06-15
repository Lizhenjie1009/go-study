package main

import "fmt"

/*
	函数变量作用域：
		全局变量：定义在函数外部的变量, 常驻内存--污染全局
		局部变量：定义在函数内部的变量, 不会常驻内存--不会污染全局
		闭包：常驻内存--不会污染全局
*/
var a = "全局变量"

func run() {
	b := "局部变量"
	fmt.Println("run", a, b)
}

// 定义函数类型
type calc func(int, int) int

// 定义int类型
type myInt int

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func main() {
	run()
	fmt.Println(a)

	var c calc // c的类型是自定义的calc函数类型
	c = add
	fmt.Printf("c-type: %T\n", c) // c-type: main.calc
	fmt.Println(c(3, 2))

	b := sub                      // 类型推断
	fmt.Printf("b-type: %T\n", b) // b-type: func(int, int) int
	fmt.Println(b(10, 5))

	var d int = 10
	var e myInt = 10
	fmt.Printf("d-type: %T\n", d) // d-type: int
	fmt.Printf("e-type: %T\n", e) // e-type: main.myInt
	fmt.Println(d + int(e))       // 相加要做类型转换
}
