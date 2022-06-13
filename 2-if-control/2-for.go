package main

import "fmt"

func main() {
	// for循环 0-10
	for i := 0; i <= 10; i++ {
		// fmt.Println(i)
	}

	i := 11
	for ; i <= 13; i++ {
		// fmt.Println(i)
	}

	j := 1
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 无限循环
	k := 5
	for {
		if k <= 10 {
			fmt.Println(k)
		} else {
			break
		}
		k++
	}

	/*
		for 嵌套
			****
			****
			****
	*/
	var row, col = 3, 4
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	/*
		for 嵌套
			*
			**
			***
	*/

	var line = 3
	for i := 0; i <= line; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	// 乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			// fmt.Print(j, "*", i, "=", i*j, "\t")
			fmt.Printf("%v*%v=%v\t", j, i, j*i)
		}
		fmt.Println("")
	}
}
