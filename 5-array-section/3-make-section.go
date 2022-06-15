package main

import "fmt"

/*
	make 函数声明切片
*/
func main() {
	// 类型，长度len，容量cap
	var a = make([]int, 5, 5)
	fmt.Println(a)                        // [0 0 0 0 0]
	fmt.Printf("%v %v\n", len(a), cap(a)) // 5	5

	// 切片修改
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	fmt.Println(a) // [10 20 30 40 50]

	b := []string{"php", "java", "node"}
	b[2] = "go"
	fmt.Println(b) // [php java go]

	// append
	var c []int
	// c[0] = 10 // 不能通过下标给切片扩容
	c = append(c, 10)
	c = append(c, 20)
	fmt.Println(c, len(c), cap(c)) // [10 20] 2 2
	c = append(c, 30, 40, 50)
	fmt.Println(c) // [10 20 30 40 50]

	// 合并切片
	d := []int{1, 2, 3}
	d = append(d, c...)
	fmt.Println(d) // [1 2 3 10 20 30 40 50]

	// 切片扩容策略
	var e []int
	for i := 0; i < 5; i++ {
		e = append(e, i)
	}
	fmt.Println(e, len(e), cap(e)) // [0 1 2 3 4] 5 8
}
