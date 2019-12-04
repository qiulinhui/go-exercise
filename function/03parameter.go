package function

import "fmt"

// Parameter2 可变参数
func Parameter2() {
	/*
		可变参数：
			概念：一个函数的参数的类型确定，但是个数不确定，就可以使用可变参数。
			语法：
				参数名 ...参数类型
				对于函数，可变参数相当于一个切片
				调用函数的时候，可以传入0-多个参数
				Println(),Printf(),Print() append() 都是可变参数的函数
			注意事项：
				A:如果一个函数的参数是可变参数，同时还有其他参数，可变参数要放在参数列表的最后
				B:一个函数的参数列表中最多只能有一个可变参数
	*/

	//求和
	getSum3()
	getSum3(1, 2, 3, 4, 5)
	getSum3(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	//切片
	s1 := []int{1, 2, 3, 4, 5}
	getSum3(s1...)
}

// 定义一个函数：用于求1-10的和
func getSum3(nums ...int) {
	sum := 0
	for i := 0; i <= len(nums); i++ {
		sum += nums[i]
	}
	fmt.Printf("总和是：%d\n", sum)
}

//一个函数中可变参数只能存在一个
func func2(s1, s2 string, nums ...float64) {

}
