package structs

import (
	"fmt"
	"strconv"
)

// Type 类型演示
func Type() {
	/*
		type 用于类型定义和类型别名
		1.类型定义：type 类型名 Type
		2.类型别名：type 类型名 = Type
	*/
	var i1 myint
	var i2 = 100
	i1 = 200
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name, s1)

	fmt.Println("-------------------------------")
	// i1 = i2 //cannot use i2 (variable of type int) as myint value in assignment
	// name = s1 //cannot use s1 (variable of type string) as mystr value in assignment
	fmt.Printf("%T,%T,%T,%T\n", i1, i2, name, s1) //structs.myint,int,structs.mystr,string
	res1 := fun1()
	fmt.Println(res1(10, 20))

	fmt.Println("-------------------------------")

	var i3 myint2
	i3 = 1000
	fmt.Println(i3)
	i3 = i2
	fmt.Println(i3)
	fmt.Printf("%T,%T,%T\n", i1, i2, i3)

}

//1.定义一个新类型
type myint int
type mystr string

//2.定义函数类型
type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

//3.类型别名
type myint2 = int // 不是重新定义新的数据类型，只是给 int 起别名，和 int 可以通用

// type MyDuration = time.Duration

// func (m MyDuration) SimpleSet() { //cannot define new methods on non-local type time.Duration

// }
