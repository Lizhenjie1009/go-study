package main

import "fmt"

func main() {
	/*
		bool声明boolean类型，值只有true、false
			默认为false
			不允许将整型转为bool
			不能参与数值运算，不能与其他类型转换
	*/

	a1 := true
	var a2 bool = false
	fmt.Printf("a1=%v type=%T\n", a1, a1)
	fmt.Println(a2)

	var a3 bool
	if a3 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
