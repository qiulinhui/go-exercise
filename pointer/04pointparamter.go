package pointer

import "fmt"

// Pparamter 指针作为参数
func Pparamter() {
	/*
		指针作为参数：
			参数的传递：值传递，引用传递
	*/
	a := 10
	fmt.Println("fun4()函数调用前，a：", a)
}

func fun4(num int) { //值传递： num = a = 10
	fmt.Println("fun4()函数中，num的值：", num)
	num = 100
	fmt.Println("fun4()函数修改num：", num)
}

func fun5(p1 *int) { //传递的是 a 的地址，就是引用传递
	fmt.Println("fun2()函数中，p1：", *p1)
	*p1 = 100
	fmt.Println("fun2()函数中，修改p1：", *p1)
}
