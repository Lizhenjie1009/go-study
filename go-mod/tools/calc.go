package tools

import (
	"fmt"
)

/*
	首字母大写--公有
	首字母小写--私有
*/

var Bbb = 20 // 公有变量

func Mul(x, y int) int { // 公有方法
	return x * y
}

func init() {
	fmt.Println("tools中的init方法")
}
