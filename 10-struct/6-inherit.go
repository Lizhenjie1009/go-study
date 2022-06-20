package main

import "fmt"

/*
	继承 inherit
		一个结构体中可以嵌套包含另一个结构体或结构体指针
*/

// 父结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Println(a.Name, "在运动")
}

// 子结构体
type Dog struct {
	Age    int
	Animal // 匿名结构体嵌套 -- 继承
}

func (d Dog) speak() {
	fmt.Println(d.Name, "汪汪汪")
}

func main() {
	var d = Dog{
		Age: 8,
		Animal: Animal{
			Name: "旺财",
		},
	}
	d.run()
	d.speak()
	fmt.Println(d)

	var d1 = Dog1{ // 嵌套结构体指针
		Age: 10,
		Animal1: &Animal1{
			Name: "阿奇",
		},
	}
	d1.run1()
	d1.speak1()
	fmt.Println(d1, *d1.Animal1)
}

// 父结构体
type Animal1 struct {
	Name string
}

func (a Animal1) run1() {
	fmt.Println(a.Name, "在运动")
}

// 子结构体
type Dog1 struct {
	Age      int
	*Animal1 // 结构体指针嵌套 -- 继承
}

func (d Dog1) speak1() {
	fmt.Println(d.Name, "汪汪汪")
}
