package main

import "fmt"

/*
	=
	+=
	-=
	*=
	/=
	%=
*/

func main() {
	var a, b = 1, 2
	a += 1
	b -= 1
	a *= 5
	a /= 2
	a %= 2            // 5 - (5/2)*2
	fmt.Println(a, b) // 1 1

	// a,b值交换 -- 中间变量
	// a,b值交换 -- 不使用中间变量，拿到和再减

	// 位运算符 - 整数二进制中操作
	var c = 5           // 101
	var d = 2           // 010
	fmt.Println(c & d)  // 000 0
	fmt.Println(c | d)  // 111 7
	fmt.Println(c ^ d)  // 111 7
	fmt.Println(c << d) // 10100 20
	fmt.Println(c >> d) // 1
}
