package main

import (
	"fmt"
	"go-mod/calc" // 引入本地的包, 只能访问公有变量
	"go-mod/tools"
	T "go-mod/tools" // 包别名
	_ "go-mod/tools" // 匿名导入包 -- 未使用不会被自动删除
)

/*
	go-mod：是一个项目，只能有一个main方法。
		go mod init 项目名   -->  	项目名要和文件名对应

	go中的 init - main: 先执行init方法在执行main方法
		在main中引入calc-init中引入tools-init，最后引入的先执行init方法
*/

func init() {
	fmt.Println("先执行init的方法")
}

func main() {
	a := calc.Add(10, 20)
	fmt.Println(a)
	fmt.Println(calc.Sub(10, 5), calc.Bbb)

	fmt.Println(tools.Mul(10, 20))
	tools.SortInt()
	tools.PrintInfo()

	T.PrintInfo()
}
