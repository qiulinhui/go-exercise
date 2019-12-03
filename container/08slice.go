package container

import "fmt"

// Slice3 切片3
func Slice3() {
	/*
		slice := arr[start,end]
			切片中的数据：[start,end]
			arr[:end],从头到 end
			arr[start:]，从 start 到尾

		从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组
			长度是从 start 到 end 切割的数据量
			但是容量是从 start 到数组的末尾

	*/

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("--------------------已有数组直接创建切片--------------------")
	s1 := a[:5]  //1-5
	s2 := a[3:8] // 4-8
	s3 := a[5:]  // 6-10
	s4 := a[:]   // 1-10

	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	// 指向同一个内存地址
	fmt.Printf("%p\n", &a) // 0xc0000a0000
	fmt.Printf("%p\n", s1) // 0xc0000a0000

	fmt.Println("---------------------长度和容量-------------------------")
	fmt.Printf("s1 len:%d,cap:%d\n", len(s1), cap(s1)) //s1 len:5,cap:10
	fmt.Printf("s2 len:%d,cap:%d\n", len(s2), cap(s2)) //s1 len:5,cap:7
	fmt.Printf("s3 len:%d,cap:%d\n", len(s3), cap(s3)) //s1 len:5,cap:5
	fmt.Printf("s4 len:%d,cap:%d\n", len(s4), cap(s4)) //s1 len:10,cap:10

	fmt.Println("---------------------更改数组的内容-------------------------")
	a[4] = 100
	fmt.Println(a)  // [1 2 3 4 100 6 7 8 9 10]
	fmt.Println(s1) // [1 2 3 4 100]
	fmt.Println(s2) // [4 100 6 7 8]
	fmt.Println(s3) // [6 7 8 9 10]
	// 引用到索引为 4 的切片都改变了

	fmt.Println("---------------------更改切片的内容-------------------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println(a)  // [1 2 3 4 100 1 1 1 1 10]
	fmt.Println(s1) // [1 2 3 4 100 1 1 1 1]
	fmt.Println(s2) // [4 100 1 1 1]
	fmt.Println(s3) // [1 1 1 1 10]
	// 改变一个切片 与其共享内存的其他切片也改变了，未共享到的索引位置不改变。

	fmt.Println("--------------------------添加元素切片扩容----------------------")
	fmt.Println(len(s1), cap(s1)) // 9 10

	s1 = append(s1, 2, 2, 2, 2, 2)
	fmt.Println(a)                // [1 2 3 4 100 1 1 1 1 10]
	fmt.Println(s1)               // [1 2 3 4 100 1 1 1 1 2 2 2 2 2]
	fmt.Println(s2)               // [4 100 1 1 1]
	fmt.Println(s3)               // [1 1 1 1 10]
	fmt.Println(len(s1), cap(s1)) // 14 20
	fmt.Printf("%p\n", s1)        // 0xc00007a000
	fmt.Printf("%p\n", &a)        // 0xc00001e050
	// 切片扩容相当于引用了另一个底层数组，内存地址发生改变。共享内存的其它切片不会一起扩容
}
