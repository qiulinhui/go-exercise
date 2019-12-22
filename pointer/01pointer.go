package pointer

import "fmt"

// Pointer 指针
func Pointer() {
	/*
		指针：pointer
			存储了另一个变量的内存地址的变量。
	*/
	//1. 定义一个int类型的变量
	a := 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("a的地址是：%p\n", &a) //0xc00009a010

	//2.创建一个指针变量，用于存储变量 a 的地址
	var p1 *int
	fmt.Println(p1)                             // <nil> 空指针
	p1 = &a                                     //p1指向了 a 的内存地址
	fmt.Println("p1 的数值是：", p1)                 // p1 中存储的是 a 的地址
	fmt.Printf("p1 自己的地址：%p\n", &p1)            // p1 中存储的是 a 的地址
	fmt.Println("p1 的值是 a 的地址，该地址存储的数据是：", *p1) //获取指针指向的变量的数值

	//3.操作变量，更改数值并不会更改内存地址
	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n", &a)

	//4.通过指针，改变变量的数值
	*p1 = 200
	fmt.Println(a)

	//5.指针的指针
	var p2 **int
	fmt.Println(p2)
	p2 = &p1
	fmt.Printf("%T,%T,%T\n", a, p1, p2) //int,*int,**int
	fmt.Println("p2的数值：", p2)           //p1的地址
	fmt.Printf("p2自己的地址：%p\n", &p2)
	fmt.Println("p2中存储的地址，对应的数值，就是p1的地址，对应的数据", *p2)
	fmt.Println("p2中存储的地址，对应的数值，再获取对应的数值", **p2)
}
