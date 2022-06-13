package main

import "fmt"

func main() {
	var a = 10

	if a > 11 {
		fmt.Println("a>5")
	} else if a < 8 {
		fmt.Println("a<20")
	} else {
		fmt.Println("no match")
	}

	if age := 30; age > 20 { // age 局部变量
		fmt.Println("大于20岁")
		fmt.Println(age)
	}
	// fmt.Println(age)

}
