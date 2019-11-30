package basic

import "fmt"

// Break 示例
func Break() {
	/*
		switch 中的 break 和 fallthrough 语句
		break:可以使用在 switch 中，也可以使用在 for 循环中
		强制结束 case 语句，从而结束 switch 分支
		fallthrough:用于穿透 switch
			当 switch 中某个 case 匹配成功之后，就执行该 case 语句
			如果遇到 fallthrough，那么后面紧邻的 case，无需匹配，直接穿透执行。
			fallthrough 应该位于某个 case 的最后一行
	*/

	n := 2
	switch n {
	case 1:
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
		fmt.Println("我是熊大")
	case 2:
		fmt.Println("我是熊二")
		fmt.Println("我是熊二")
		break // 用于强制结束 case，意味着 switch 被强制结束
		// fmt.Println("我是熊二") //break 后面的语句不会执行
	case 3:
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
		fmt.Println("我是光头强")
	}

	fmt.Println("--------------------")
	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	}
}
