package basic

import "fmt"

func if2() {
	/*
		if 语句的其他写法：
		if 初始化语句；条件 {
			//初始化的变量会随着 if 语句的结束而销毁
		}
	*/

	if num := 4; num > 0 {
		fmt.Println("正数", num)
	} else if num < 0 {
		fmt.Println("负数", num)
	}
	// fmt.Println(num) //undefined:num

	num2 := 5
	if num2 > 0 {
		fmt.Println("num2是正数", num2)
	}
	fmt.Println(num2)
}
