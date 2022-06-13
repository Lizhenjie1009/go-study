package main

import "fmt"

/*
	Println: 换行，空格
	Print: 不换行，无空格

	Printf: 格式化
*/
func main() {
	// fmt.Println("你好 gogogo")
	// fmt.Println("hi")
	// fmt.Print("hello")

	// var a = "aaa" // 声明必须要使用
	// fmt.Println("hi", a)
	// fmt.Printf("%v", a)

	/*
		占位符
			%v: 输出一个变量
			%b: 输出二进制
			%o: 输出八进制
			%d: 输出十进制
	*/

	var c int = 1
	var d int = 2
	var e int = 3
	fmt.Println("c=", c, "d=", d, "e=", e)
	fmt.Printf("c=%v d=%v e=%v\n", c, d, e)
	fmt.Printf("c=%v d=%v e=%v", c, d, e)

	// 类型推导
	f := 10
	g := 20
	h := 30
	fmt.Println(f, g, h)
	fmt.Printf("f=%o g=%o h=%o", f, g, h)

	// 使用Printf打印一个变量的类型
	fmt.Printf("\nf=%v f的类型是%T", f, f)
}
