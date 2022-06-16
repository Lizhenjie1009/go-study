package main

import "fmt"

/*
	1.结构体匿名字段: 类型不能重复
	2.结构体的字段类型: 基本数据类型, 切片, map, 结构体

	注意: 结构体字段类型是: 指针, 切片, map的零值都是nil, 即还未分配空间
			需要先make
*/

// 结构体匿名字段
type Person struct {
	string
	int
	// int
}

type Body struct {
	height int
	sex    string
	Hobby  []string
	map1   map[string]string
}

func (b *Body) setHobby(hobby string) {
	b.Hobby[2] = hobby
}

func main() {
	p := Person{
		"zs",
		20,
	}
	fmt.Println(p)

	b := Body{
		height: 180,
		sex:    "男",
	}
	// b.Hobby = append(b.Hobby, "singe", "paly")
	b.Hobby = make([]string, 3, 6)
	b.Hobby[0] = "single"
	b.Hobby[1] = "dance"
	b.Hobby[2] = "play"
	b.map1 = make(map[string]string)
	b.map1["name"] = "ll"
	b.map1["home"] = "bj"
	fmt.Printf("%#v\n", b)
	b.setHobby("do")
	fmt.Printf("%v\n", b.Hobby)
}
