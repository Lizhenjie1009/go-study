package main

import (
	"fmt"
	"reflect"
)

/*
	结构体反射获取结构体的属性
	t.Elem(): 基本类型可以直接拿到所对应元素的key或value
	t.Elem().FieldByName("Name"): 和基本类型不一致的是，结构体需要通过反射方法先拿到所对应元素的key或value
*/

type Student struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() (str string) {
	str = fmt.Sprintf("姓名:%v, 年龄:%v, 成绩:%v", s.Name, s.Age, s.Score)
	fmt.Println(str)
	return
}

func (s *Student) SetInfo(name string, age, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Printf("打印...%#v", s)
}

func main() {
	s1 := Student{
		Name:  "小明",
		Age:   18,
		Score: 80,
	}
	s1.GetInfo()
	s1.SetInfo("小红", 19, 90)
	s1.GetInfo()

	// PrintStructField(s1)
	PrintStructFn(&s1) // 修改传入指针地址
}

// 反射获取结构体的字段
func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// 当前类型和原始类型都不是结构体
	if v.Kind() != reflect.Struct && v.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体")
		return
	}

	// 1.通过***类型变量(TypeOf)***里面的Field可以获取结构体字段
	t0 := t.Field(0)        // 结构体
	fmt.Printf("%#v\n", t0) // reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4e5b40), Tag:"json:\"name\"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Printf("字段名称:%v, 字段类型:%v, 字段tag:%v\n", t0.Name, t0.Type, t0.Tag.Get("form"))
	fmt.Println(t0.Tag)

	fmt.Println("--------------------------------------------")

	// 2.通过类型变量里面的FieldByName可以获取结构体字段
	t1, ok := t.FieldByName("Age")
	if ok {
		fmt.Printf("字段名称:%v, 字段类型:%v, 字段tag:%v\n", t1.Name, t1.Type, t1.Tag.Get("json"))
	}

	// 3.通过类型变量里面的NumField可以获取结构体有几个字段
	fieldCount := t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性Type") // 3

	fmt.Println("--------------------------------------------")

	// 4.通过值变量获取结构体属性对应的值
	v0 := v.Field(0) // 结构体对应的value
	v1 := v.FieldByName("Name")
	fmt.Println(v0, v1) // 小红

	fmt.Println("--------------------------------------------")

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, v.Field(i), v.FieldByName(t.Field(i).Name), t.Field(i).Tag, t.Field(i).Type, t.Field(i))
	}
}

// 反射执行结构体的方法
func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}
	fmt.Println(t, v)
	// 1.通过类型字面量里面的Method可以获取结构体的方法
	m0 := t.Method(0)                 // index和结构体方法夫人顺序没有关系，和结构体方法的ASCII码有关系
	fmt.Println(m0.Name, m0.Type, m0) // GetInfo   func(main.Student) string   {GetInfo  func(main.Student) string <func(main.Student) string Value> 0}
	fmt.Println("--------------------------------------------")

	// 2.通过类型变量获取结构体有多少个方法
	num := t.NumMethod()
	fmt.Println(num, "个方法")
	fmt.Println("--------------------------------------------")

	m1, ok := t.MethodByName("GetInfo")
	if ok {
		fmt.Println(m1.Name) // GetInfo
		fmt.Println(m1.Type) // func(main.Student) string
	}
	fmt.Println("--------------------------------------------")

	// 3.通过《值变量》执行方法(注意需要使用值变量，并且要注意参数) v.Method(0).Call(nil) 或者 v.MethodByName(fnName).Call(nil)
	// v.Method(1).Call(nil)
	// v.MethodByName("Print").Call(nil)
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)
	fmt.Println("--------------------------------------------")

	// 4.执行方法传入参数(注意需要使用《值变量》, 并且要注意参数，接收的参数是[]reflect.Value的切片)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("小兰"))
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(100))
	v.MethodByName("SetInfo").Call(params) // 执行方法传入参数
	v.MethodByName("GetInfo").Call(nil)

	// 5.执行方法获取方法的值
}
