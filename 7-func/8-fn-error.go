package main

import (
	"errors"
	"fmt"
)

/*
	异常：
		panic：抛出异常
		recover：拦截异常
*/

func test1() {
	fmt.Println("开始-test")
}
func test2() {
	defer func() {
		// 没有异常err为nil
		err := recover()
		if err != nil { // 有异常
			fmt.Println(err, "recover 拦截异常")
		}
	}()
	panic("panic: throw new error 终止下面的代码执行-报错")
}

func main() {
	test1()
	test2()
	fmt.Println("结束-test")

	doFile()
	fmt.Println("go")
}

func doFile() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(e)
		}
	}()

	err := readFile("x.go")
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("操作文件失败")
	}
}
