package main

import "fmt"

/*
	切片 section
		切片的本质就是对底层数组的封装
*/
func main() {
	// 声明切片
	var n []int // 默认值是  nil
	fmt.Println(n == nil)

	var a = []int{1, 2, 3}
	a = append(a, 4)
	fmt.Println(a)

	for k, v := range a {
		fmt.Println(k, v)
	}

	// 基于数组定义切片
	b := [5]int{55, 56, 57, 58, 59}       // 数组
	c := b[:]                             // 切片	[:] 获取数组中的所有值
	fmt.Printf("%v c=%T b=%T\n", c, c, b) // [55 56 57 58 59] c=[]int b=[5]int
	// 截取数组 - 根据索引 左包 右不包
	d := b[1:4]
	e := b[1:]           // 从索引1截取到最后
	f := b[:4]           // 从索引4之前
	fmt.Println(d, e, f) // [56 57 58]	[56 57 58 59]		[55 56 57 58]

	// 基于切片定义切片 - 切片再切片
	g := []string{"bj", "sh", "gz", "sz", "cd", "cq"}
	i := g[1:]
	fmt.Println(g, i)

	/*
		切片的长度及容量
			长度：元素的个数
			容量：第一个元素到其底层数组元素末尾的个数
	*/
	h := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度: %v 容量: %v\n", len(h), cap(h))
	fmt.Printf("长度: %d 容量: %d\n", len(h), cap(h)) // 长度: 5 容量: 5
	j := h[2:]
	fmt.Printf("长度: %d 容量: %d\n", len(j), cap(j)) // 长度: 4 容量: 4
	k := h[1:3]                                   // 3, 5   容量为5: 从元素3到底层数组元素末尾
	fmt.Printf("长度: %d 容量: %d\n", len(k), cap(k)) // 长度: 2 容量: 5
	l := h[:3]                                    // 2, 3, 5	容量为5: 从元素2到底层数组元素末尾
	fmt.Printf("长度: %d 容量: %d\n", len(l), cap(l)) // 长度: 3 容量: 6

}
