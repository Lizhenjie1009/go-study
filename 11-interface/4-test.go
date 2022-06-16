package main

import "fmt"

/*
	结构体指针接收参数
	结构体实现多个接口
*/

type Animaler interface {
	setName(string)
	getName() string
}

type Animaler2 interface {
	run()
}

type Dog struct {
	Name string
}

func (d *Dog) setName(name string) {
	d.Name = name
	fmt.Println(d)
}
func (d *Dog) getName() string {
	return d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) setName(name string) {
	c.Name = name
	fmt.Println(c)
}
func (c Cat) getName() string {
	return c.Name
}
func (c Cat) run() {
	fmt.Println(c.Name, "run")
}

func main() {
	d := &Dog{
		Name: "旺财",
	}
	var d1 Animaler = d
	d1.setName("阿刚")

	c := &Cat{
		Name: "小花",
	}
	var c1 Animaler = c
	fmt.Println(c1.getName())
	c1.setName("如花")

	// 结构体实现多个接口
	cc := &Cat{
		Name: "小花",
	}
	var cc1 Animaler = cc
	var cc2 Animaler2 = cc
	cc1.setName("多接口的cat")
	cc2.run()
}
