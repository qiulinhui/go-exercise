package function

import "fmt"

// Parameter3 参数传递
func Parameter3() {
	/*
		参数传递：
			A:值传递：传递的是数据的副本。修改数据，对于原始的数据没有影响。
				值传递类型的数据，默认都是值传递：基础类型，array，struct
			B:引用传递：传递的是数据的地址
				引用类型的数据，默认都是引用传递：slice，map，chan

	*/

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前数组的数据：", arr1)
	fun1(arr1)
	fmt.Println("函数调用后数组的数据：", arr1)

	fmt.Println("-------------------------------------")
	s1 := []int{1, 2, 3, 4}
	fmt.Println("函数调用前切片的数据：", s1)
	fun2(s1)
	fmt.Println("函数调用后切片的数据：", s1)
}

func fun1(arr2 [4]int) {
	fmt.Println("函数中，数组的数据：", arr2)
	arr2[0] = 100
	fmt.Println("函数中，数组的数据修改后：", arr2)
}

func fun2(s2 []int) {
	fmt.Println("函数中，切片的数据：", s2)
	s2[0] = 100
	fmt.Println("函数中，切片的数据修改后：", s2)
}
