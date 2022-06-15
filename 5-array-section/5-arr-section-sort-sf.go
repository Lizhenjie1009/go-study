package main

import (
	"fmt"
	"sort"
)

func main() {
	// [1,3,5,7,8] 找出和为8的两个元素下标索引 -> (0,3) (1,2)
	var arr = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ { // i+1 从i往后加, 避免重复
			if arr[i]+arr[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

	/*
		选择排序：从小到大排序  [9,8,7,6,5,4]
		假设：
		外层循环 -- 内层循环 => 结果
		9 -- 8,7,6,5,4 => 8,7,6,5,4,9
		8 -- 7,6,5,4,9 => 7,6,5,4,8,9
											6,5,4,7,8,9
											5,4,6,7,8,9
		5 -- 4,6,7,8,9 => 4,5,6,7,8,9
	*/
	var q = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(q); i++ {
		for j := i + 1; j < len(q); j++ {
			if q[i] > q[j] {
				// 中间变量-数据交换
				temp := q[i]
				q[i] = q[j]
				q[j] = temp
			}
		}
	}
	fmt.Println(q) // [4 5 6 7 8 9]

	/*
		冒泡排序：每一轮找到最大的值放在最后面
			每一轮比完最大的在最后面
			依次减去每一轮循环后的值
	*/
	var m = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m)-1-i; j++ {
			if m[j] > m[j+1] {
				temp := m[j]
				m[j] = m[j+1]
				m[j+1] = temp
			}
		}
	}
	fmt.Println(m) // [4 5 6 7 8 9]

	a := []int{2, 3, 4, 1, 7, 6, 5}
	b := []string{"b", "d", "c", "a"}
	c := []float64{5.5, 2.2, 4.4, 1.1, 3.3}
	// 升序
	sort.Ints(a)
	sort.Strings(b)
	sort.Float64s(c)
	fmt.Println(a, c, b)

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sort.Sort(sort.Reverse(sort.StringSlice(b)))
	sort.Sort(sort.Reverse(sort.Float64Slice(c)))
	fmt.Println(a, c, b)

}
