package main

import "fmt"

func main() {

	// switch case
	var ext = ".java"
	switch ext {
	case ".html":
		fmt.Println("1")
		break
	case ".css":
		fmt.Println("2")
		break
	case ".js":
		fmt.Println("3")
		break
	default:
		fmt.Println(0)
	}

	switch num := 1; num {
	case 1, 2, 3:
		fmt.Println(true)
	default:
		fmt.Println(false)
	}

	var age = 23
	switch {
	case age < 24:
		fmt.Println("hhxx")
	case age > 30:
		fmt.Println("money")
	}

	// switch穿透 fallthrough 兼容c语言的case设计
	var scope = 65
	switch { // B C
	case scope >= 70:
		fmt.Println("A")
	case scope < 70 || scope >= 60:
		fmt.Println("B")
		fallthrough
	default:
		fmt.Println("C")
	}
}
