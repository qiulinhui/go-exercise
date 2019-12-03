package container

import "fmt"

// Slice4 切片与数据类型
func Slice4() {
	/*
		按照类型来分：
			基本类型：int, float, string, bool
			复合类型：array, slice, map, struct, pointer, function, chan
		按照特点来分：
			值类型：int, float, string, bool, array
				传递的是数据的副本
			引用类型：slice
				传递的是地址，多个变量指向了同一块内存地址
		所以：切片是引用类型的数据，存储了底层数组的引用
	*/

	//1. 数组
	a1 := [4]int{1, 2, 3, 4}
	a2 := a1
	fmt.Println(a1, a2)
	a1[0] = 100
	fmt.Println(a1, a2)

	//2. 切片
	s1 := []int{1, 2, 3}
	s2 := s1
	fmt.Println(s1, s2)
	s1[0] = 100
	fmt.Println(s1, s2)

	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
}
