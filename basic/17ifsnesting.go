package basic

import "fmt"

// Ifsnesting if 语句嵌套示例
func Ifsnesting() {
	/*
		if 语句的嵌套：
			if 条件 1 {
				A段
			} else {
				if 条件 2 {
					B段
				} else {
					C段
				}
			}

			简写：
			if 条件 1 {
				A段
			} else if 条件 2 {
				B段
			} else if 条件 3 {
				C段
			}
	*/

	sex := "未知"
	if sex == "男" {
		fmt.Println("可以去男厕所")
	} else if sex == "女" {
		fmt.Println("可以去女厕所")
	} else {
		fmt.Println("迷茫")
	}
	fmt.Println("函数结束")

}
