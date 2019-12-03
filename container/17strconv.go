package container

import (
	"fmt"
	"strconv"
)

// Strconv strconv包的使用
func Strconv() {
	/*
		strconv包：字符串和基本数据类型之间的转换
			string convert
	*/
	//1.bool类型
	s1 := "true"
	//Parse 字符串转其他类型
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1)
	// Format 其他类型转字符串
	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n", ss1, ss1)

	//2.整数
	s2 := "100"
	// 参数：源，进制，位数
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2)

	ss2 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T,%s\n", ss2, ss2)

	//itoa() atoi()
	i3, err := strconv.Atoi("-42") //转换为 int 类型
	fmt.Printf("%T,%d\n", i3, i3)
	ss3 := strconv.Itoa(-42)
	fmt.Printf("%T,%s\n", ss3, ss3)
}
