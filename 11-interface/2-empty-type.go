package main

import "fmt"

/*
	空接口

	类型断言
		x.(T)
			x: 类型为interface{}空接口的变量
			T: 断言x可能是的类型
*/

type Empty interface{} // 表示没有任何约束, 任意的类型都可以实现

// 空接口作为函数的参数，可以传任意类型
func add(a, b interface{}) {
	fmt.Printf("a:%v, t:%T; b:%v, t:%T\n", a, a, b, b)
}

func main() {
	var a Empty
	str := "go"
	a = str
	fmt.Println(a)

	num := 20
	a = num
	fmt.Println(a)

	var b interface{} // b是空接口类型，可以是任意类型
	b = 10
	b = "20"
	b = true
	fmt.Println(b)

	add(12, "12")
	add("12", 12)

	var p = make(map[string]interface{}) // map的值可以是任意类型
	p["name"] = "zs"
	p["age"] = 20
	p["married"] = false
	fmt.Println(p, p["name"])

	var q = []interface{}{1, 2, "a", "b", true} // 切片可以是任意类型
	fmt.Println(q)

	/*
		类型断言
		x.(T)
			x: 类型为interface{}空接口的变量
			T: 断言x可能是的类型
	*/
	v1, ok := q[0].(int)
	v2, o2 := q[2].(string)
	v3, o3 := q[4].(string)
	v3t, o3t := q[4].(bool)
	fmt.Println(v1, ok)
	fmt.Println(v2, o2)
	fmt.Println(v3, o3)
	fmt.Println(v3t, o3t)

	v4, o4 := p["married"].(bool)
	fmt.Println(v4, o4)

	myPrint(10)
	myPrint("10")
	myPrint(true)
	myPrint([]int{1, 2})

	myPrint1(10)
	myPrint1("10")
	myPrint1(true)
	myPrint1([]int{1, 2})
}

// 传入任意的类型，实现不同的功能
func myPrint(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println(v, "string")
	} else if v, ok := x.(int); ok {
		fmt.Println(v, "int")
	} else if v, ok := x.(bool); ok {
		fmt.Println(v, "bool")
	} else {
		fmt.Println("no match")
	}
}

// x.(type) 判断一个变量的类型, 只能用在switch语句
func myPrint1(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("no match")
	}
}
