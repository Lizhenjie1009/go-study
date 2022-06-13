package main

import (
	"fmt"
)

/*
	切片是引用数据类型
*/
func main() {
	// 引用同一份地址
	var a = []int{1, 2, 3, 4}
	b := a
	a[0] = 0
	fmt.Println(a, b) // [0 2 3 4] [0 2 3 4]

	// copy() 复制切片
	c := make([]int, len(a), cap(a))
	copy(c, a)
	c[0] = 1
	fmt.Println(a, c) // [0 2 3 4] [1 2 3 4]

	// 从切片中删除元素
	d := []int{1, 2, 3, 4, 5}
	d = append(d[:2], d[3:]...) // [1,2] [4,5] 使用append合并方式进行删除
	fmt.Println(d)              // [1 2 4 5]

}
