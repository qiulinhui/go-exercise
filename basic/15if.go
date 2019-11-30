package basic

import "fmt"

// If if示例
func If() {
	/*
		条件语句 if
		语法格式：
		if 条件表达式 {
			//
		}
	*/

	// 1. 给定一个数字，如果大于 10 ，就输出打印这个数字大于 10

	num := 6
	if num > 10 {
		fmt.Println("大于10")
	}

	//2. 给定一个成绩，如果大于等于 60 分，就打印及格
	score := 88
	if score >= 60 {
		fmt.Println(score, "成绩及格...")
	}
	fmt.Println("函数结束")

}
