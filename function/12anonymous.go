package function

import "fmt"

//Anonymous 匿名函数的使用
func Anonymous() {
	/*
		匿名函数：没有名字的函数
		定义一个匿名函数，直接进行调用。通常只能使用一次。也可以使用匿名函数赋值给某个函数变量，那么就可以调用多次了。

		匿名函数：
			Go语言是支持函数式编程：
			1.将匿名函数作为另一个函数的参数，回调函数
			2.将匿名函数作为另一个函数的返回值，可以形成闭包结构
	*/

	func() {
		fmt.Println("我是一个匿名函数")
	}() // 函数后面加小括号 直接调用
	func16 := func() { // 将匿名函数赋值给函数变量
		fmt.Println("我也是一个匿名函数")
	}
	func16()
	func16() // 将一个匿名函数赋值给一个函数变量可以多次调用

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) //匿名函数调用了，将执行结果赋值给了res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //将匿名函数赋值给了 res2

	fmt.Println(res2)
	fmt.Println(res2(100, 200))
}
