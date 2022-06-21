package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score int
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

func main() {
	std := Student{
		Name:  "小明",
		Age:   20,
		Score: 98,
	}

	reflectChangeStructKey(&std) // 修改结构体，传入地址

	std.GetInfo()
}

// 修改结构体属性的值
func reflectChangeStructKey(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr {
		fmt.Println("不是指针类型, 需要结构体指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体类型, 需要结构体指针类型")
		return
	}

	name := v.Elem().FieldByName("Name")
	name.SetString("小花")
	age := v.Elem().FieldByName("Age")
	age.SetInt(21)
}
