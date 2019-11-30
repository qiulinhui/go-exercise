package basic

import "fmt"

// Switch2 switch演示2
func Switch2() {
	/*
		1. switch的标准写法：
			switch 变量 {
			case 数值1:
				分支1
			case 数值2:
				分支2
			default:
				默认分支
			}
		2. 省略 switch 后的变量，相当于直接作用在 true 上
			switch { //true
			case ture :
			case false :
			}
		3. case 后可以同时跟随多个数值
			switch 变量{
			case 数值1,数值2,数值3:分支1
			case 数值4,数值5,数值6:分支2
			}
		4. switch 后面可以多一条初始化语句
			switch 初始化语句; 变量 {

			}
	*/
	switch {
	case true:
		fmt.Println("true...")
	case false:
		fmt.Println("false...")
	}

	score := 88
	switch {
	case score >= 0 && score < 60:
		fmt.Println(score, "不及格")
	case score >= 60 && score < 70:
		fmt.Println(score, "及格")
	case score >= 70 && score < 80:
		fmt.Println(score, "中等")
	case score >= 80 && score < 90:
		fmt.Println(score, "良好")
	case score >= 90 && score < 100:
		fmt.Println(score, "优秀")
	default:
		fmt.Println("成绩有误。。。")
	}
	letter := ""
	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println(letter, "是元音")
	case "M", "N":
		fmt.Println("M或者N")
	default:
		fmt.Println("其他")
	}

	switch language := "golang"; language {
	case "golang":
		fmt.Println("Go语言。。。")
	case "java":
		fmt.Println("Java语言。。")
	case "python":
		fmt.Println("python语言。。")
	}
}
