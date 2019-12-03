package container

import "fmt"

// Slice5 深拷贝和浅拷贝
func Slice5() {
	/*
		深拷贝：拷贝的是数据本身
			值类型的数据，默认都是深拷贝：array, int, float, bool, struct
		浅拷贝：拷贝的是数据地址
			导致多个变量指向同一块内存
			引用类型的数据，默认都是浅拷贝：slice, map

			因为切片是引用类型的数据，直接拷贝的是地址
	*/

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) //len:0,cap:0
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	// copy() 专门用于切片深拷贝
	s3 := []int{7, 8, 9}
	fmt.Println(s2)
	fmt.Println(s3)

	copy(s2, s3)         //将 s3 中的元素，拷贝到 s2 中
	copy(s3, s2)         //将 s2 中的元素，拷贝到 s3 中
	copy(s3[1:], s2[2:]) //拷贝一部分数据
	fmt.Println(s2)
	fmt.Println(s3)

}
