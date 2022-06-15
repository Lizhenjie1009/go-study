package main

import "fmt"

/*
	defer: 延迟执行
		先defer后执行，后defer先执行

	提示：***defer注册要延迟执行的函数时，该函数所有的参数都要确定其值***
*/

func fn1() {
	fmt.Println("开始-fn1")
	defer func() {
		fmt.Println("func-1")
		fmt.Println("func-2")
	}()
	fmt.Println("结束-fn1")
}

// 匿名返回值
func fn2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

// 命名返回值
func fn3() (a int) {
	defer func() {
		a++
	}()
	return a
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	// fmt.Println("开始")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("结束")
	// fn1()

	// fmt.Println(fn2()) // 0
	// fmt.Println(fn3()) // 1

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	/*
		A 1 2 3
		B 10 2 12
		BB 10 12 22
		AA 1 3 4
	*/

}
