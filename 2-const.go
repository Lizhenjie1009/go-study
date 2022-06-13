package main

import "fmt"

func main() {
	var name = "zs"
	name = "ls"

	const pi = 3.14
	const (
		key = "key"
		num = 10
	)
	fmt.Println(name, pi, key, num)

	// const声明 iota 从const第一个值开始初始化为0 (自增长)
	const (
		a = 5
		b
		c = iota
		d
	)

	fmt.Println(a, b, c, d)
}
