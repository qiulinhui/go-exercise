package basic

import "fmt"

// For2 for循环的其他写法
func For2() {
	/*
		1. 标准写法：
			for 表达式1; 表达式2; 表达式3 {
				循环体
			}
		2. 同时省略表达式1和表达式3
			for 表达式2 {

			}
			相当于 while (条件)
		3. 同时省略3个表达式
			for {

			}
			相当于 while(true)
			注意点：当 for 循环中，省略了表达式 2，就相当于直接作用在了 ture 上
		4. 其他的写法：for 循环中同时省略几个表达式都可以。。
		省略表达式1：
		省略表达式2：循环永远成立-->死循环
		省略表达式3：
	*/

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-->", i)
}
