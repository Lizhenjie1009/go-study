package main

import "fmt"

// 引用数据类型必须初始化才能赋值
func main() {
	// var user map[string]string
	// user["name"] = "zs"
	// fmt.Println(user) // error 引用数据类型, 没有分配内存空间

	var user = make(map[string]string)
	user["name"] = "zs"
	fmt.Println(user)

	// var a *int
	// *a = 100
	// fmt.Println(a)  // error 指针也是引用数据类型, 没有分配内存空间

	// new 指针类型初始化 - 给指针变量分配储存空间
	var a = new(int) // a是一个指针变量, 类型*int, 值0
	fmt.Printf("值:%v, 类型:%T, 地址:%p\n", a, a, &a)
	*a = 100 // 改变指针所对应变量的值
	fmt.Println(*a)

	// make函数分配内存, 只用于切片/map/channel, 返回原类型
	// new函数分配内存, 返回指针
	var f = new(bool)
	fmt.Println(f, *f)
}
