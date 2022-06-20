package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = 'a'                      // 字符可以用单引号
	fmt.Printf("%v type=%T\n", a, a) // 97 int32
	// %c 原样输出
	fmt.Printf("%c type=%T\n", a, a) // a int32

	/*
		uint8
		rune
	*/

	var str = "this"
	fmt.Printf("%c %v type=%T\n", str[2], str[2], str[2]) // 105 uint8

	// 一个汉字(utf-8)3个字节，一个字母(ascii)1个字节
	// unsafe.Sizeof 查看不了string类型所占用的储存空间
	fmt.Println(len(str), unsafe.Sizeof(str)) // 4 16
	var str1 = "你好go"
	fmt.Println(len(str1)) // 8

	// 汉字字符使用的是utf-8
	var str2 = '国'
	fmt.Println(str2)                 // 22269
	fmt.Printf("%v %c\n", str2, str2) // 22269 国

	// 循环输出字符
	str3 := "你好 golang"
	for i := 0; i < len(str3); i++ { // byte
		fmt.Printf("%v (%c)", str3[i], str3[i])
	}
	fmt.Println()
	for _, r := range str3 { // rune 以utf-8输出
		fmt.Printf("%v (%c)", r, r)
	}
	fmt.Println()

	// 修改字符串的字符 先转为rune类型的切片或byte类型的切片
	str4 := "big"
	byteStr := []byte(str4)
	byteStr[0] = 'p'
	fmt.Println(str4, string(byteStr))

	str5 := "中国china"
	runeStr := []rune(str5)
	runeStr[0] = '大'
	runeStr[2] = '大'
	fmt.Println(str5, string(runeStr))
}
