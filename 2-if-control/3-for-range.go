package main

import "fmt"

/*
	for range 遍历数组、切片、字符串、map、及通道(channel)
		数组、切片、字符串返回索引和值
		map返回键和值
		通道(channel)返回通道内的值
*/

func main() {
	var str = "你好 go"
	for k, v := range str {

		// fmt.Println(k, v) // v=字符对应的字节码
		fmt.Printf("k=%v v=%c\n", k, v)
	}

	// 切片
	var arr = []string{"a", "b", "c"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
	for k, v := range arr {
		fmt.Println(k, v)
	}

}
