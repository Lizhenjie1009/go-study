package main

import "fmt"

// break			跳出当前循环
// continue		跳出这次循环
// goto				无条件跳转, 快速跳出
func main() {
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	// label(xxx) 跳出对应的循环
xxx:
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break xxx
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}

	// continue
	for i := 0; i < 10; i++ {
		if i > 2 && i < 8 {
			continue
		}
		fmt.Println(i)
	}

yyy:
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue yyy
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}

	// goto
	var n = 30
	if n > 24 {
		fmt.Println("成年人")
		goto label1
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("ccc")
label1:
	fmt.Println("ddd")
}
