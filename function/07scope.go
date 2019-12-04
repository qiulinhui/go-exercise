package function

import "fmt"

// num3 := 1000 //全局变量不支持简短定义的写法
var num3 = 1000

// Scope 函数的作用域
func Scope() {
	/*
		作用域：变量可以使用的范围
			局部变量：函数内部定义的变量，就叫做局部变量。
					函数在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了。
			全局变量：函数外部定义的变量叫做全局变量。
					所有的函数都可以使用，而且共享这一份数据
	*/

	// 定义在 Scope 函数中，所以 n 的作用域就是 Scope 函数的范围内
	n := 10
	fmt.Println(n)
	if a := 1; a <= 10 {
		fmt.Println(a)
		fmt.Println(n)
	}
	// fmt.Println(a) //不能访问 a, 出了作用域，a 的作用域只在 if 语句内
	fmt.Println(n)
	if b := 1; b <= 10 {
		n := 20
		fmt.Println(b) //1
		fmt.Println(n) //20,就近原则
	}

	fun6()
	fun7()
	fmt.Println("Scope中访问全局变量：", num3) //2000
}

func fun6() {
	//fmt.Println(n)
	num1 := 100
	fmt.Println("fun1()函数中 num1 的值：", num1)
	num3 = 2000
	fmt.Println("fun1()函数，访问全局变量：", num3)
}

func fun7() {
	//fmt.Println(n)
	num1 := 200
	fmt.Println("fun1()函数中 num1 的值：", num1)
	fmt.Println("fun1()函数，访问全局变量：", num3)
}
