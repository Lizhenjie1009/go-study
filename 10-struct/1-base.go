package main

import "fmt"

/*
	结构体 - struct
		type关键字定义结构体
*/

// 自定义类型
type myInt int
type myFn func(int, int) int

// 类型别名
type myFloat = float64

func add(x, y int) int {
	return x + y
}

/*
	人的结构体 person struct
		name age sex
*/
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var a myInt = 10
	fmt.Printf("v=%v t=%T\n", a, a) // v=10 t=main.myInt

	var b myFn
	b = add
	fmt.Printf("v=%v t=%T\n", b(1, 2), b) // v=3 t=main.f1

	var c myFloat = 3.2
	fmt.Printf("%v %f %T\n", c, c, c) // 3.2 3.200000 float64

	var p Person // 实例化Person结构体
	p.name = "zs"
	p.age = 30
	p.sex = "男"
	fmt.Printf("值:%v 类型:%T\n", p, p)  // 值:{zs 30 男} 类型:main.Person
	fmt.Printf("值:%#v 类型:%T\n", p, p) // %#v:打印详细信息	 值:main.Person{name:"zs", age:30, sex:"男"} 类型:main.Person
}
