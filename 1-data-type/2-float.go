package main

import (
	"fmt"
	"unsafe"
)

/*
	float32  4个字节
		3.4e38 = 3.4乘10的38次方
	float64  8个字节
		1.8e308 = 1.8乘10的308次方
*/
func main() {
	var a1 float64 = 3.12

	fmt.Printf("%v -- %f -- %T\n", a1, a1, a1)
	fmt.Println(unsafe.Sizeof(a1))

	// %f 输出float类型    %.2f输出保留2位小数
	var a2 float64 = 3.54159266589
	fmt.Printf("%v %f %.2f\n", a2, a2, a2)
	fmt.Printf("%.4f\n", a2)

	// go 默认浮点数是 float64
	a3 := 3.14
	fmt.Printf("%T\n", a3)

	// 科学计数法表示浮点类型

	var a4 float32 = 3.14e2 // 3.14 * 10的2次方   314
	fmt.Printf("a4=%v float=%f type=%T\n", a4, a4, a4)

	var a5 float32 = 3.14e-2 // 3.14 / 10的2次方   0.0314
	fmt.Printf("a5=%v float=%f type=%T\n", a5, a5, a5)

	// 精度丢失
	var a6 float64 = 1129.6
	fmt.Println(a6 * 100) // 112960 => 112959.99999999999
	a7 := 8.2
	a8 := 3.8
	fmt.Println(a7 - a8) // 4.4 => 4.3999999999999995

	// 类型转换 int转float
	a9 := 10
	b1 := float64(a9)
	// c1 := int(b1)
	fmt.Printf("a9 type %T\n", a9)
	fmt.Printf("b1 type %T\n", b1)
	// fmt.Printf("c1 type %T\n", c1)
}
