package main

import (
	"fmt"
	"time"
)

// %02d 整数不够2位，往前补0
func main() {
	timeObj := time.Now()
	fmt.Println(timeObj) // 2022-06-14 14:59:39.0127538 +0800 CST m=+0.001634001
	y := timeObj.Year()
	m := timeObj.Month()
	d := timeObj.Day()
	h := timeObj.Hour()
	mm := timeObj.Minute()
	ss := timeObj.Second()

	fmt.Println(y, m, d, h, mm, ss)
	fmt.Printf("%v-%v-%v %v:%v:%v\n", y, m, d, h, mm, ss)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", y, m, d, h, mm, ss)

	/*
		2006 年
		01 月
		02 日
		03 时 12小时制  --  15 24小时制
		04 分
		05 秒
	*/
	f := timeObj.Format("2006-01-02 03:04:05")
	f1 := timeObj.Format("2006/01/02 15:04:05")
	fmt.Println(f, f1)

	// 获取当前的时间戳
	// unixTime := timeObj.Unix() // 获取当前时间戳
	unixTime := time.Now().Unix()       // 毫秒
	unixNaTime := time.Now().UnixNano() // 纳秒时间戳
	fmt.Println(unixTime, unixNaTime)

	// 时间戳转日期字符串
	date := time.Unix(time.Now().Unix(), 0) // 2022-06-14 15:31:30 +0800 CST
	delTime := date.Format("2006-01-02 15:04:05")
	fmt.Println(delTime)

	// 日期字符串转时间戳
	str := "2022-06-14 15:34:00"
	fm := "2006-01-02 15:04:05"
	a, err := time.ParseInLocation(fm, str, time.Local)
	fmt.Println(a, err)   // 2022-06-14 15:34:00 +0800 CST <nil>
	fmt.Println(a.Unix()) // 1655192040
}
