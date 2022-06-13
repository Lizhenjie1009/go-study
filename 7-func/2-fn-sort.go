package main

import (
	"fmt"
	"sort"
)

// 选择排序 -- 升序true、降序false -- 是否克隆 -- 命名返回值
func chooseSort(data []int, toggle, deep bool) (res []int) {
	if deep {
		res = make([]int, len(data), cap(data))
		copy(res, data)
	} else {
		res = data
	}
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if toggle && res[i] > res[j] {
				temp := res[i]
				res[i] = res[j]
				res[j] = temp
			} else if res[i] < res[j] {
				temp := res[i]
				res[i] = res[j]
				res[j] = temp
			}
		}
	}
	return
}

/*
	m1 := map[string]string{
		"name": "zs",
		"age": "20",
		"sex": "男",
		"height": "180",
	}
	根据key排序
*/
// map				--				升序true、降序false		 -- 		命名返回值
func mapSort(data map[string]string, toggle bool) (str string) {
	var qp []string
	for k, _ := range data {
		qp = append(qp, k)
	}
	if toggle {
		sort.Strings(qp)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(qp)))
	}

	for _, v := range qp {
		// str += v + "=>" + data[v] + "\n"
		str += fmt.Sprintf("%v: %v\n", v, data[v])
	}
	return
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	fmt.Println(a, chooseSort(a, false, true))

	m1 := map[string]string{
		"name":   "zs",
		"age":    "20",
		"sex":    "男",
		"height": "180",
	}
	fmt.Println(mapSort(m1, false))
}
