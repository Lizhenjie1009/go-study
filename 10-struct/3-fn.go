package main

import "fmt"

/*
	结构体是值类型
*/

type Person struct {
	name   string
	age    int
	sex    string
	height int
}

// 	(变量 结构体类型)
func (a Person) PrintInfo() string {
	fmt.Printf("姓名: %v, 年龄: %v\n", a.name, a.age)
	return a.name
}

// (变量 指针结构体类型) --> 结构体是值类型, 所以用指针
func (a *Person) setInfo(name string, age int) {
	a.name = name
	a.age = age
}

// 自定义类型添加方法
type myInt int

func (a myInt) pInfo() {
	fmt.Println(a, "hhhh")
}

func main() {
	var p = Person{
		name: "zs",
		age:  20,
		sex:  "男",
	}
	p1 := p
	p1.name = "ll"
	fmt.Printf("p: %#v\n", p)   // main.Person{name:"zs", age:20, sex:"男"}
	fmt.Printf("p1: %#v\n", p1) // main.Person{name:"ll", age:20, sex:"男"}

	fmt.Println(p.PrintInfo())
	fmt.Println(p1.PrintInfo())

	p.setInfo("王五", 30)
	fmt.Println(p)

	var a myInt = 20
	a.pInfo()
}
