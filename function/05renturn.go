package function

import "fmt"

// Return 函数的返回值
func Return() {
	/*
		函数的返回值：
			一个函数的执行结果，返回给函数的调用处。执行结果就叫做函数的返回值。
		return语句：
			一个函数的定义上有返回值，那么函数中必须使用 return 语句，将结果返回给调用处。
			函数返回的结果必须和函数返回的一致：类型，个数，顺序
			1.将函数的结果返回给调用处
			2.同时结束了该函数的执行
		空白标识符，专门用来舍弃数据：_
	*/

	//1.设计一个函数，用于求1-10的和，将结果在主函数中打印输出
	res := getSum4()
	fmt.Println("1-10的和：", res)

	fmt.Println(getSum5()) //5050

	res1, res2 := rectangle(5, 3)
	fmt.Println("周长：", res1, "面积：", res2)
	res3, res4 := rectangle2(5, 3)
	fmt.Println("周长：", res3, "面积：", res4)
}

// 定义一个函数，带返回值
func getSum4() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

// 定义函数时，指明要返回的数据是哪个
func getSum5() (sum int) {
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return
}

//go支持多返回值函数
//定义函数用于求矩形的周长和面积
func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

// 定义函数返回值时可以预定义返回的变量名
func rectangle2(len, wid float64) (perimeter, area float64) {
	perimeter = (len + wid) * 2
	area = len * wid
	return
}
