package main

import "fmt"

/*
	结构体值接收者和指针接收者实现接口的区别
		值接收者：实例化后结构体值类型和结构体指针类型都可以赋值给接口变量
		指针接收者：实例化后结构体指针类型可以赋值给接口变量
*/

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

// 值接收者
func (p Phone) start() {
	fmt.Println(p.Name, "启动了")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机了")
}

func main() {
	p := Phone{
		Name: "华为手机",
	}
	var p1 Usber = p // 值类型传参 -- 实例化后的值类型结构体实现Usber接口
	p1.start()

	p2 := &Phone{
		Name: "小米手机",
	}
	var p3 Usber = p2 // 值类型传参 -- 实例化后的指针类型结构体实现Usber接口
	p3.start()

	c := &Camera{
		Name: "佳能",
	}
	var c1 Usber = c // 指针类型传参 -- 实例化后的指针类型结构体实现Usber接口
	c1.start()
}

type Camera struct {
	Name string
}

// 指针类型
func (c *Camera) start() {
	fmt.Println(c.Name, "开机")
}
func (c *Camera) stop() {
	fmt.Println(c.Name, "关机")
}
