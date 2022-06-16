package calc

import (
	"fmt"
	_ "go-mod/tools"
)

/*
	首字母大写--公有
	首字母小写--私有
*/

var aaa = 10 // 私有变量
var Bbb = 20 // 公有变量

func Add(x, y int) int { // 公有方法
	return x + y + aaa
}

func Sub(x, y int) int { // 公有方法
	return x - y
}

func init() {
	fmt.Println("calc中的init方法")
}
