package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	// num := 10 + x // x是空接口类型，不能运算

	// 1使用类型断言
	v, _ := x.(int)
	num := 10 + v
	fmt.Println(num)

	// 2反射reflect.ValueOf返回的是reflect.value
	v1 := reflect.ValueOf(x)
	// n := v1 + 10 // reflect.value类型不一致不能运算
	n := 10 + v1.Int() // v1.Int()获取v1的原始值
	fmt.Println(n)
}

func main() {
	// 使用reflect.ValueOf的值
	reflectValue(10)

	// 判断reflect.ValueOf的值
	var a float32 = 1.23
	var b float64 = 1.23
	reflectValue1(a)
	reflectValue1(b)
	reflectValue1(10)
	reflectValue1("go")
	reflectValue1(true)
}

func reflectValue1(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Int: // 获取的reflect对应的类型
		fmt.Println("int类型原始值", v.Int())
	case reflect.Float32:
		fmt.Println("Float32类型原始值", v.Float())
	case reflect.Float64:
		fmt.Println("Float64类型原始值", v.Float())
	case reflect.String:
		fmt.Println("String类型原始值", v.String())
	default:
		fmt.Println("no match", x)
	}

}
