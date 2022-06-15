package main

import "fmt"

// 数组是值类型 - 基本数据类型和数组都是值类型
// []中没有长度或者 ... 就是切片, 切片是引用类型
func main() {
	var arr [3]int
	var arr1 [4]int
	var str [5]string
	fmt.Printf("arr %v %T\n", arr, arr)
	fmt.Printf("arr1 %v %T\n", arr1, arr1)
	fmt.Printf("str %v %T\n", str, str)

	var a = []int{1, 2, 3, 4, 5} // 切片
	fmt.Println(a, len(a), a[0])

	var b [3]int
	b[0] = 11
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	fmt.Println(c)

	var d [5]int
	for k, _ := range d {
		d[k] = k
	}
	fmt.Println(d)

	e := []int{1: 10, 2: 20, 5: 50} // 切片
	fmt.Println(len(e), e)

	var f = []int{1, 2, 3} // 切片是引用类型
	g := f
	f[0] = 11
	fmt.Println(f, g)

	// 多维数组
	// var h = [3][2]int{
	var h = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(h, h[0][0])

	for _, w := range h {
		for _, v := range w {
			fmt.Println(v)
		}
	}
}
