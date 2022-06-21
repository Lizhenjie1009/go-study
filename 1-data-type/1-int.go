package main

import (
	"fmt"
	"unsafe"
)

/*
	基本数据类型
		整形(int)、浮点型(float64)、布尔型(bool)、字符串(string)

	复合数据类型
		数组(array)、切片(slice)、结构体(struct)、函数(func)、map、通道(chan)、接口(interface)等

	指针(ptr)
*/

// 整形: 有符号整型，无符号整型

/*
	有符号整型长度: int8、int16、int32、int64
	对应无符号整型: uint8、uint16、uint32、uint64
*/
func main() {
	// int
	var num int = 10
	num = 20
	fmt.Printf("%v %T\n", num, num)

	// int8的范围 负2的七次方~2的七次方-1
	// -128 ~ 127
	var num8 int8 = -128
	fmt.Printf("%v\n", num8)

	// unsafe.Sizeof 查看不同长度的整型，在内存里面的存储空间
	// int8是一个字节，1024byte=1kb
	fmt.Printf("num8=%v 类型%T\n", num8, num8)
	fmt.Println(unsafe.Sizeof(num8), "字节")

	// uint8的范围 0~2的八次方
	// 0~255
	var n8 uint8 = 255
	fmt.Printf("%v 类型%T\n", n8, n8)
	fmt.Println(unsafe.Sizeof(n8), "字节")

	// int类型转换
	// var a1 int32 = 32
	// var a2 int64 = 64
	var (
		a1 int32 = 10
		a2 int64 = 21
	)
	// fmt.Println(a1 + a2)
	fmt.Println(int64(a1) + a2) // 把a1转换成64位

	// 高位向地位转换
	var a3 int16 = 130
	fmt.Println(int8(a3)) // -126 溢出

	/*
		数字字面量语法
			%d：表示10进制输出
			%b：表示2进制输出
			%o：表示8进制输出
			%x：表示16进制输出
	*/

	a4 := 30
	fmt.Printf("a4=%x\n", a4)
	fmt.Println(unsafe.Sizeof(a4)) // int64
}
