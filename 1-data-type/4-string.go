package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string = "hi go"
	var s2 = "hi go1"
	s3 := "hi go3"

	fmt.Println(s1, s2, s3)
	fmt.Printf("s2-type-%T\n", s2)

	/*
		string转义字符
			\r 回车符-返回行首
			\n 换行符
			\t 制表符
			\' 单引号
			\" 双引号
			\\ 反斜杠
	*/
	s4 := "hi\tgo\n\\\"helloworld\"\\"
	fmt.Printf("ab\tcd\n")
	fmt.Println(s4)

	// 多行字符串 ``
	s5 := `
		1
			2
				3
			4
		5
	`
	fmt.Println(s5)

	// len 字符长度
	fmt.Println(len(s5))

	// 拼接字符串：+ 、fmt.Sprintf
	fmt.Println(s4 + s3)
	s6 := fmt.Sprintf("%v - %v", s4, s3)
	fmt.Println(s6)

	// 分割字符--切片
	var s7 string = "123-456-789"
	arr := strings.Split(s7, "-")
	fmt.Println(arr) // [123 456 789] 切片

	// 把切片连接成字符串
	s8 := strings.Join(arr, "*")
	fmt.Println(s8)

	// 切片
	qp := []string{"js", "vue", "react"}
	fmt.Println(qp)
	fmt.Printf("qp=type=%T\n", qp)

	// 字符串是否包含
	var s9 = "this is str"
	var s10 = "this"
	s11 := strings.Contains(s9, s10)
	fmt.Println(s11)

	// 前缀或者后缀
	s12 := strings.HasPrefix(s9, s10)
	s13 := strings.HasSuffix(s9, "str")
	fmt.Println(s12, s13)

	// 字符串出现的位置, 找不到返回 -1
	s14 := strings.Index(s9, "is")
	s15 := strings.LastIndex(s9, "s")
	fmt.Println(s14, s15) // 2 8
}
