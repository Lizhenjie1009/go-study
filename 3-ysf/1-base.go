package main

import "fmt"

/*
	+ - * / %
*/

func main() {
	var a, b = 6, 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// 除法注意：结果为整数相除去掉小数，保留整数
	// 浮点数相除会保留小数
	var c, d = 10, 3
	fmt.Println(c / d)    // 3
	fmt.Println(10.0 / 3) // 3.3333333333333335

	// 取余注意：余数=被除数-(被除数/除数)*除数
	fmt.Println(10 % 3)  // 1
	fmt.Println(-10 % 3) // -1  = -10 - (-10 / 3) * 3
	fmt.Println(10 % -3) // 1  = 10 - (10 / -3) * -3

	// 单独使用 ++ --
	c--
	d++
	fmt.Println(c, d)
}
