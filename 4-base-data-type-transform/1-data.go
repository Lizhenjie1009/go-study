package main

import "fmt"

/*
	数值转换
	数据类型转换 低转高
	整型和浮点型转换，转成浮点型 -> 浮点型转整型会丢失精度
*/

func main() {
	var (
		a int8    = 20
		b int16   = 40
		c float32 = 20.203
		d float64 = 40
	)

	fmt.Println(int16(a) + b)
	fmt.Println(float64(c) + d)
	fmt.Println(int16(c) + b) // 丢失精度
	fmt.Println(c + float32(b))

}
