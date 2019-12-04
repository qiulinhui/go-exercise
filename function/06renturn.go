package function

import "fmt"

// Return2 return的作用
func Return2() {
	/*
		return语句：词义"返回"
			A：一个函数有返回值，那么使用 return 将返回值返回给调用处
			B：同时意味着结束了函数的执行
		注意点：
			1.一个函数定义了返回值，必须使用 return 语句将结果返回给调用处。return 后的数据必须和函数定义的一致：个数，类型，顺序。
			2.可以使用_，来舍弃多余的返回值
			3.如果一个函数定义了有返回值，那么函数中有分支，循环，那么要保证，无论执行了哪个分支，都要有 return 语句被执行到
			4.如果一个函数没有定义返回值，那么函数中也可以使用 return，专门用于结束函数的执行
	*/

	a, b, c := fun3()
	fmt.Println(a, b, c)

	// 舍弃部分返回值
	_, _, d := fun3()
	fmt.Println(d)
	fmt.Println(fun4(-30))
}

func fun3() (int, float64, string) {
	return 0, 0, "hello"
}

func fun4(age int) int {
	if age >= 0 {
		return age
	}
	// } else {
	fmt.Println("年龄不能为负。。")
	return 0
	// }
}

// return 用于结束函数的执行
func fun5() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			// break  //终止循环
			return //终止函数
		}
		fmt.Println(i)
	}
	fmt.Println("func5()...over...")
}
