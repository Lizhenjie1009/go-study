package main

import "fmt"

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

type calcType func(int, int) int

// 函数作为另一个函数的参数
func calc(x, y int, cb calcType) int {
	return cb(x, y)
}

// 函数作为返回值
// func do(op string) func(int, int) int {
func do(op string) calcType {
	switch op {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	fmt.Println(calc(10, 5, add))
	fmt.Println(calc(10, 5, sub))
	// 传一个匿名函数
	fmt.Println(calc(10, 5, func(x, y int) int {
		return x * y
	}))

	fmt.Println(do("+")(10, 5))
	fmt.Println(do("-")(10, 5))
	fmt.Println(do("*")(10, 5))

	// 匿名函数
	func() {
		fmt.Println("自执行匿名函数...")
	}()
	a := func(x, y string) string {
		return x + " -- " + y
	}
	fmt.Println(a("a", "b"))

	var x, y = 1, 2
	func(x, y int) {
		fmt.Println(x + y)
	}(x, y)
}
