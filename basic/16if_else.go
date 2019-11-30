package basic

import "fmt"

// IfElse if else 示例
func IfElse() {
	/*
		if ... else 语句
			if 条件 {
				//条件成立，执行此处的代码
				A 段
			} else {
				// 条件不成立，执行此处的代码
				B 段
			}
	*/

	//给定一个成绩，如果大于等于 60 分，就打印及格
	score := 44
	if score >= 60 {
		fmt.Println(score, "成绩及格...")
	} else {
		fmt.Println(score, "成绩不及格...")
	}

	// 给定性别如果是男，就去男厕所，否则去女厕所
	sex := "男"
	if sex == "男" {
		fmt.Println("可以去男厕所")
	} else {
		fmt.Println("可以去女厕所")
	}
	fmt.Println("函数结束")

}
