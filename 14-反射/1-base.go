package main

import (
	"fmt"
	"reflect"
)

/*
	反射：
		反射是指程序运行期间对程序本身进行访问和修改的能力。

	反射使用场景：
		反射可以在判断空接口保存数据的类型和值是什么的时候，也可以用类型断言
		反射可以把结构体序列化成json字符串，自定义结构体Tab标签
		ORM框架

	go中的变量：
		类型信息：预先定义好的元信息。
		值信息：程序运行过程中可动态变化的。

	go的反射机制，任何接口值都是由一个***具体类型***和***具体类型的值***两部分组成的。

	反射相关功能内置的reflect：
		任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成。
		reflect包提供了reflect.TypeOf和reflect.ValueOf来获取任意对象的Type和Value。

	反射类型划分为两种：类型(Type)和类(Kind)
		类(Kind)：底层的类型
*/

// 反射获取任意变量的类型
func main() {
	reflectFn(10)
	reflectFn(10.123)
	reflectFn("10")
	reflectFn(true)
	reflectFn([2]int{1, 2})
	reflectFn([]string{"a", "b", "c"})
	reflectFn(map[string]string{"name": "a", "name1": "b", "name2": "c"})
	reflectFn(make(chan int, 3))

	var a myInt = 123
	reflectFn(a)
	p := Person{
		Name: "zs",
		Age:  20,
	}
	reflectFn(p)
	reflectFn(p.Age)
	var b = 10
	reflectFn(&b) // *int
}

type myInt int
type Person struct {
	Name string
	Age  int
}

func reflectFn(x interface{}) {
	t := reflect.TypeOf(x)  // 反射获取空接口的任意类型
	v := reflect.ValueOf(x) // 反射获取空接口的任意类型的值
	n := t.Name()           // 反射获取类型的名称
	k := t.Kind()           // 反射获取类型的种类
	fmt.Printf("类型种类:%v,\t 类型名称:%v,\t 类型:%v,\t 类型值:%v, 值:%v;\n", k, n, t, v, x)
}
