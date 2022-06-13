package main

import "fmt"

func main() {
	// make 创建 map 类型
	var user = make(map[string]string)
	user["name"] = "zs"
	user["age"] = "20"
	user["sex"] = "男"
	fmt.Println(user, user["name"])

	for k, v := range user {
		fmt.Println(k, v)
	}

	// 声明填充
	// map[key-type][value-type]
	person := map[string]string{
		"name": "ls",
		"age":  "30",
		"sex":  "男",
	}
	fmt.Println(person)

	// map类型数据的curd
	var p1 = make(map[string]string)
	p1["name"] = "ww"
	p1["age"] = "40"
	// 创建-修改
	p2 := map[string]string{
		"name": "ww",
	}
	p2["name"] = "aa"
	fmt.Println(p1, p2["name"])

	// 查找 v值 t是否存在
	v, t := p2["name"]
	v1, t1 := p2["xxx"]
	fmt.Println(v, t)   // aa true
	fmt.Println(v1, t1) //    false

	// 删除key和value
	p3 := map[string]string{
		"name":   "ls",
		"age":    "30",
		"sex":    "男",
		"height": "180",
	}
	delete(p3, "age")
	fmt.Println(p3)
}
