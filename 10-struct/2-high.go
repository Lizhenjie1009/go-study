package main

import "fmt"

/*
	结构体 struct
		输出事物的全部属性或部分, 就需要用到自定义数据类型, 可以封装多个基本数据类型
		结构体变量名小写表示私有的
*/

type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 Person // 初始化
	p1.name = "zs"
	p1.age = 20
	p1.sex = "男"
	fmt.Printf("值: %#v, t: %T\n", p1, p1) // 值: main.Person{name:"zs", age:20, sex:"男"}, t: main.Person

	var p = new(Person) //  对应指针
	p.name = "ls"
	p.age = 30
	p.sex = "男"
	(*p).name = "ll"
	fmt.Printf("值: %#v, t: %T\n", p, p) // 值: &main.Person{name:"ll", age:30, sex:"男"}, t: *main.Person

	var p2 = &Person{} //  对应指针
	p2.name = "ww"
	p2.age = 40
	p2.sex = "男"
	fmt.Printf("值: %#v, t: %T\n", p2, p2) // 值: &main.Person{name:"ww", age:40, sex:"男"}, t: *main.Person

	var p3 = Person{ // 初始化 + 实例化
		name: "lk",
		age:  50,
		sex:  "女",
	}
	fmt.Printf("值: %#v, t: %T\n", p3, p3) // 值: main.Person{name:"lk", age:50, sex:"女"}, t: main.Person

	var p4 = &Person{ // 对应指针 + 实例化
		name: "lp",
		age:  60,
		sex:  "女",
	}
	fmt.Printf("值: %#v, t: %T\n", p4, p4) // 值: &main.Person{name:"lp", age:60, sex:"女"}, t: *main.Person

	var p5 = &Person{ // 对应指针 + 实例化 + 值映射一一对应
		"lo",
		70,
		"女",
	}
	fmt.Printf("值: %#v, t: %T\n", p5, p5) // 值: &main.Person{name:"lo", age:70, sex:"女"}, t: *main.Person
}
