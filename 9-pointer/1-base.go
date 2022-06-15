package main

import "fmt"

/*
	指针 pointer
		指向另一个变量内存地址
		指针也是变量, 特殊的变量
*/

// %p -> 内存地址
// &变量名 -> 取内存地址
// *p 表示取出p这个变量对应变量内存地址的值
func main() {
	var a = 10
	fmt.Printf("a: %v; type: %T; adde: %p\n", a, a, &a) // a: 10; type: int; adde: 0xc000014098

	var p *int = &a // 指针变量 	类型 *int
	// var p = &a
	fmt.Printf("p: %v; type: %T; adde: %p\n", p, p, &p) //p: 0xc000014098; type: *int; adde: 0xc000006030

	// 根据p打印a的值
	fmt.Println(*p) // 10 -> *p 表示取出p这个变量对应变量内存地址的值

	*p = 20
	fmt.Println(p, *p, a) //0xc000014098 20 20

	var b = 5
	fn1(b)
	fmt.Println(b) // 5
	fn2(&b)
	fmt.Println(b) // 20
}

func fn1(x int) {
	x = 10
}
func fn2(x *int) {
	*x = 20
}
