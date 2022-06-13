package main

import (
	"fmt"
)

func main() {
	fmt.Println("go 声明变量")
	/*
		1.var
			var 变量名
			var 变量名 类型
	*/
	var num = 10
	var str string = "str"

	fmt.Printf("%v", str)
	fmt.Printf("%v\n", num)

	/*
		2.定义多个变量
			var 变量名,变量名 类型
			var (
				变量名 类型
				变量名 类型
			)
	*/

	var a1, a2 string
	a1 = "a1"
	a2 = "a2"

	var (
		b1 string
		b2 int
	)
	b1 = "b1"
	b2 = 10

	var (
		c1 = "c1"
		c2 = 20
	)
	fmt.Println(a1, a2, b1, b2, c1, c2)

	/*
		3. := 短变量声明
			只能声明局部变量，不能用于全局变量
	*/

	d1 := "d1"
	d2 := 30

	e1, e2, e3 := "e1", 40, 3.14
	fmt.Println(d1, d2, e1, e2, e3)
	fmt.Printf("%T\n", e3)

	/*
		4.匿名变量(anonymous variable)，用于忽略某个值
			用下划线表示
	*/
	// var username, age = getUserInfo()
	var name, _ = getUserInfo()
	// var _, age = getUserInfo()
	fmt.Println(name)
}

func getUserInfo() (string, int) {
	return "张三", 30
}
