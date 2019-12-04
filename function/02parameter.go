package function

import "fmt"

// Parameter 函数的参数
func Parameter() {
	/*
		函数的参数：
			形式参数：也叫形参。函数定义的时候，用于接收外部传入的数据的变量
				函数中，某些变量的数值无法确定，需要由外部传入数据。
			实际参数：也叫实参。函数调用的时候，给形参赋值的实际数据
		函数调用：
			1.函数名：声明的函数名和调用的函数名要统一
			2.实参必须严格匹配形参：顺序，个数，类型必须一一对应
	*/
	//求1-10的和
	getSum2(10)
	//求1-20的和
	getSum2(20)
	//求1-100的和
	getSum2(100)

	getAdd(20, 10)
	getAdd2(1, 2)
	func1(1.3, 1.4, "hello")

}

// 定义一个函数：用于求1-10的和
func getSum2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和是：%d\n", n, sum)
}

// 多个参数
func getAdd(a int, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d \n", a, b, sum)
}

//多个参数类型相同的简便写法
func getAdd2(a, b int) {
	fmt.Printf("a:%d，b:%d \n", a, b)
}

//调用时要根据参数定义的顺序传入相应的值
func func1(a, b float64, c string) {
	fmt.Printf("a:%.2f,b:%.2f,c:%s\n", a, b, c)
}
