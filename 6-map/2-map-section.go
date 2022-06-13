package main

import (
	"fmt"
	"sort"
	"strings"
)

// map不初始化的类型是nil
// map是引用数据类型
func main() {
	// map类型的切片
	// 切片
	var p1 = []string{"zs", "ls"}
	var p2 = make([]string, 3, 3)
	fmt.Println(p1, p2)

	// map
	var p3 = make(map[string]string)
	fmt.Println(p3)

	// map-section
	var p4 = make([]map[string]string, 2, 3)
	if p4[0] == nil { // map默认值是nil，不能操作
		p4[0] = make(map[string]string)
		p4[0]["name"] = "p4-0"
	}
	if p4[1] == nil {
		p4[1] = make(map[string]string)
		p4[1]["name"] = "p4-1"
	}
	fmt.Println(p4, p4[0] == nil)
	for _, v := range p4 {
		for k, v1 := range v {
			fmt.Println(k, v1)
		}
	}

	// map类型的值为切片
	var p5 = make(map[string][]string)
	p5["sports"] = []string{"篮球", "乒乓球", "羽毛球"}
	p5["id"] = []string{"1", "2", "3"}
	fmt.Println(p5)

	// map排序
	p6 := make(map[int]int)
	p6[5] = 0
	p6[1] = 1
	p6[3] = 2
	p6[8] = 20
	for k, v := range p6 {
		fmt.Println(k, v)
	}
	/*
		5 0
		1 1
		3 2
		8 20
	*/
	// 根据key升序输出map -> 签名算法
	var p7 []int // 定义切片排序
	for k, _ := range p6 {
		p7 = append(p7, k)
	}
	sort.Ints(p7)
	fmt.Println(p7)
	for _, v := range p7 {
		fmt.Println(v, p6[v])
	}

	// 统计单词出现的次数
	var str = "how do you do"
	s := strings.Split(str, " ") // [how do you do]
	p8 := make(map[string]int)
	// for i := 0; i < len(s); i++ {
	// 	p8[s[i]] += 1
	// }
	for _, v := range s {
		p8[v]++
	}
	fmt.Println(p8) // map[do:2 how:1 you:1]
}
