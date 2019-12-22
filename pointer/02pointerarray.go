package pointer

import "fmt"

//PArray 数组指针和指针数组
func Parray() {
	/*
		数组指针：首先是一个指针，一个数组的地址。
			*[4]Type
		指针数组：首先是一个数组，存储的数据类型是指针
		[4]*Type

		前面是[就是指针数组，前面是*就是数组指针

		*[5]float64
		*[3]string
		[3]*string
		[5]*float64
		*[5]*float64
		*[3]*string
		**[4]string
		**[4]*string
	*/

	//1.创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//2.创建一个指针，存储该数组的地址-->数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1) //&[1 2 3 4]
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", &p1)

	//3.根据数组指针操作数组
	(*p1)[0] = 100
	fmt.Println(arr1)

	p1[0] = 200 //简化写法
	fmt.Println(arr1)

	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2) //[1 2 3 4]
	fmt.Println(arr3)
	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)
	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a)

	b = 1000
	fmt.Println(arr2)
	fmt.Println(arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}

}
