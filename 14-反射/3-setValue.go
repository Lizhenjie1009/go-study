package main

import (
	"fmt"
	"reflect"
)

/*
	设置反射的值
*/

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		// v.SetInt(10) // 值类型不能这么修改
		fmt.Println(v, v.Int())
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// fmt.Println(v.Kind())        // ptr
	// fmt.Println(v.Elem().Kind()) // int64  -- v.Elem()对应真实原始元素
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(10)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("golang")
	}
}

func main() {
	var a int64 = 100
	reflectSetValue1(a)
	reflectSetValue2(&a)
	fmt.Println(a)

	var b string = "gogogo"
	reflectSetValue2(&b)
	fmt.Println(b)
}
