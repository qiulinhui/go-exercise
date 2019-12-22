package function

import "fmt"

// Function2 函数的本质
func Function2() {
	/*
		Go语言的数据类型：
			数值类型：整数，浮点
				进行运算操作，加减乘除，打印
			字符串：
				可以获取单个字符，截取子串，遍历，strings包下的函数操作。。
			数组，切片，map：
				存储数据，修改数据，获取数据，遍历数据。。。
			函数：
				加()，进行调用
			注意点：
				函数作为一种复合数据类型，可以看作是一种特殊的变量。
					函数名（）：将函数进行调用，函数中的代码会全部执行，然后将return的结果返回给调用处
					函数名：指向函数体的内存地址
	*/

	//函数作一个变量
	fmt.Printf("%T\n", func14) //func(int,int)
	fmt.Println(func14)        //0x48eb90看作函数名对应的函数体的地址

	//直接定义一个函数类型的变量
	var c func(int, int)
	fmt.Println(c) //<nil> 空

	c = func14 //将func14的值（函数体的地址赋值给c）
	fmt.Println(c)
	func14(10, 20)
	c(100, 200) //c也是函数体 加上小括号就可以被调用

	res1 := func15       //将func15的值（函数的地址）赋值给res1，res1和func15指向同一个函数体
	res2 := func15(1, 2) //将func15函数进行调用，将函数的执行结果赋值给res2，相当于：a+b
	fmt.Println(res1)
	fmt.Println(res2)
}

func func14(a, b int) {
	fmt.Printf("a:%d,b:%d\n", a, b)
}

func func15(a, b int) int {
	return a + b
}
