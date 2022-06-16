package main

import "fmt"

/*
	结构体嵌套 - struct nesting
*/

type User struct {
	Username string
	Password string
	// Address  Address // user结构体嵌套address结构体
	Address        // 嵌套匿名结构体
	AddTime string // 结构体和嵌套结构体同名
}

type Address struct {
	Name    string
	Phone   string
	City    string
	AddTime string
}

func main() {
	var u User
	u.Username = "admin"
	u.Password = "123456"
	u.Address.Name = "zs"
	u.Address.Phone = "12144123133"
	u.Address.City = "北京"
	// 匿名结构体可以这样查找
	u.City = "上海" // 访问结构体成员的时候, 首先在结构体中查找, 找不到再去匿名结构体中查找
	u.AddTime = "2022-06-15"
	u.Address.AddTime = "2022-06-16"
	fmt.Printf("%#v", u)
}
