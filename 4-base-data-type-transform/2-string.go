package main

import (
	"fmt"
	"strconv"
)

/*
	其他类型转string
		fmt.Sprintf()
		strconv

	string转数值类型
*/

func main() {
	var (
		i int     = 20
		f float64 = 12.456
		t bool    = true
		b byte    = 'a'
	)
	var str = fmt.Sprintf("%d", i)
	var str2 = fmt.Sprintf("%f", f)
	var str3 = fmt.Sprintf("%t", t)
	var str4 = fmt.Sprintf("%c", b)
	fmt.Printf("%v %T\n", str, str)
	fmt.Printf("%v %T\n", str2, str2)
	fmt.Printf("%v %T\n", str3, str3)
	fmt.Printf("%v %T\n", str4, str4)

	// strconv.FormatInt(int64, 进制)
	var a1 int = 20
	str5 := strconv.FormatInt(int64(a1), 10)
	fmt.Printf("%v %T\n", str5, str5)

	// strconv.FormatFloat(float64, 格式化类型, 保留几位小数, 格式化类型64、32)
	var a2 float32 = 3.1415926
	str6 := strconv.FormatFloat(float64(a2), 'f', -1, 64)
	fmt.Printf("%v %T\n", str6, str6)

	// strconv.FormatBool(boolean)
	var a3 bool = true
	str7 := strconv.FormatBool(a3)
	fmt.Printf("%v %T\n", str7, str7)

	var a4 = 'a'
	str8 := strconv.FormatUint(uint64(a4), 10)
	fmt.Printf("%v %T\n", str8, str8)

	// string转数值类型
	var a5 = "123456"
	str9, _ := strconv.ParseInt(a5, 10, 64)
	// fmt.Printf("%v %v %T\n", err, str9, str9)
	fmt.Printf("%v %T\n", str9, str9)

	str10, _ := strconv.ParseFloat(a5, 32)
	fmt.Printf("%v %T\n", str10, str10)

}
