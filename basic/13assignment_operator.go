package basic

import "fmt"

func assignment() {
	/*
		赋值运算符：
		=，+=，-=，*=，/=，%=，<<=，>>=，&=，|=，^=
		=：把 = 右侧的数值赋值给 = 左侧的变量
		+=：a += b，相当于 a = a + b
	*/

	var a int
	a = 3
	a += 4
	fmt.Println(a) // 7

	a -= 3
	fmt.Println(a) // 4
	a *= 2
	fmt.Println(a) // 8
	a /= 3
	fmt.Println(a) // 2
	a %= 1
	fmt.Println(a) // 0
}
