package main

/*
	third party -- 引入第三方的包
		1. go mod tidy				拉取依赖
		2. go mod download 		下载依赖
		3. go mod vendor			将依赖复制到项目中
*/

import (
	"fmt"

	"github.com/shopspring/decimal" // 先写好要引入的三方包，执行上面流程
	"github.com/tidwall/gjson"
)

func main() {
	// str := "https://github.com/shopspring/decimal" // 全局 go get github.com/shopspring/decimal

	a := 3.1
	b := 3.2
	// a + b = 6.300000000000001
	c := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)) // 6.3
	fmt.Println(a+b, c)

	a1 := 8.2
	b1 := 3.8
	// a1 - b1 = 4.3999999999999995
	c1 := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)) // -0.1
	fmt.Println(a1-b1, c1)

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.first")
	fmt.Println(value)
}
