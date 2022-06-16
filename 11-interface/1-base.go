package main

import "fmt"

/*
	接口是一种抽象数据类型 -- 为了约束数据
		接口定义了对象的行为规范，只定义规范不实现。
		接口中定义的规范由具体对象来实现。
*/

type Usber interface { // 接口名后加 er 来表示是接口
	start(string, string)
	stop()
}

// 如果接口中有方法的话，必须要通过结构体或者自定义类型来实现该接口
type Phone struct {
	id   string
	Name string
}

// Phone实现usb接口的话，要实现usb接口中的所有方法
func (p Phone) start(a, b string) {
	fmt.Println(p.Name, "开机了")
}
func (p Phone) stop() {
	fmt.Println("phone stop")
}

func main() {
	p := Phone{
		id:   "das16354651312313asd1asfda",
		Name: "诺基亚",
	}
	p.start("NO", "KIA")
	p.stop()

	var p1 Usber // 接口就是一个数据类型
	p1 = p       // 表示手机实现usb接口
	fmt.Println(p1)
	p1.start("1", "2")

	c := Camera{}
	var c1 Usber = c // 表示相机实现usb接口
	c1.start("佳", "能")
	c1.stop()
	// c1.run() // 接口之外的方法，只能通过自身实例调用
	c.run(c1)
	// c.run(c)
	// c.run(p1)
	// c.run(p)
	// c.run(只要满足Usber接口类型就可以)
}

type Camera struct {
}

func (c Camera) start(a, b string) {
	fmt.Println(a, b, "开机了")
}
func (c Camera) stop() {
	fmt.Println("camera 关机了")
}
func (c Camera) run(usb Usber) {
	_, o := usb.(Camera)
	fmt.Println(o)
	usb.start("run1", "run2")
	usb.stop()
	fmt.Println("camera run")
}
