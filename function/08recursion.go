package function

import "fmt"

// Recursion 递归函数
func Recursion() {
	/*
		递归函数（recursion）：一个函数自己调用自己，就叫做递归函数。
			递归函数要有一个出口，逐渐的向出口靠近
	*/

	//1. 求1-5的和

	sum := getSum6(5)
	fmt.Println(sum)
	res := getFiboacci(12)
	fmt.Println(res)
}

func getSum6(n int) int {
	if n == 1 {
		return 1
	}
	return getSum6(n-1) + n
}

// 获取斐波那契数
func getFiboacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFiboacci(n-1) + getFiboacci(n-2)
}
