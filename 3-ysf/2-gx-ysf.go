package main

import "fmt"

func main() {
	/*
		==
		!=
		>
		<
		>=
		<=
	*/
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	flag := 1 > 2
	fmt.Println(flag)

	/*
		&&
		||
		!
	*/
	fmt.Println(4 > 2 && 3 > 2)
	fmt.Println(1 > 2 || 3 > 2)
	fmt.Println(!false)
}
