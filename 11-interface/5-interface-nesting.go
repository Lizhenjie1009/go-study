package main

import "fmt"

/*
	接口的嵌套
	空接口和类型断言的使用细节
*/

type A interface {
	setName(string)
}
type B interface {
	getName()
}
type Animaler interface { // 接口的嵌套
	A
	B
}
type Dog struct {
	Name string
}

func (d *Dog) setName(name string) {
	d.Name = name
	fmt.Println("setName", d.Name)
}
func (d Dog) getName() {
	fmt.Println("getName", d.Name)
}

func main() {
	d := &Dog{
		Name: "来福",
	}
	var d1 Animaler = d
	d1.setName("旺财")
	d1.getName()

	// 空接口和类型断言的使用细节
	var user = make(map[string]interface{})
	user["name"] = "张三"
	user["hobby"] = []string{"睡觉", "吃饭", "打豆豆"}
	fmt.Println(user["name"])
	fmt.Println(user["hobby"])
	// fmt.Println(user["hobby"][0]) // 错误：空接口不知道当前数据类型
	v, o := user["hobby"].([]string) // 使用类型断言的方式来获取空接口对应的值
	fmt.Println(v, o, v[0])

	var address = Address{
		Name:  "李四",
		Phone: 123123131412,
	}
	user["address"] = address
	// fmt.Println(user["address"].Name) // 错误：空接口不知道当前数据类型
	v1, o1 := user["address"].(Address) // 使用类型断言的方式来获取空接口对应的值
	fmt.Println(v1, o1, v1.Name)
}

type Address struct {
	Name  string
	Phone int
}
