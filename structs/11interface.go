package structs

import "fmt"

//Interface3 接口嵌套
func Interface3() {
	/*
		接口嵌套
	*/

	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("-----------")
	var a1 A1 = cat
	a1.test1()
	fmt.Println("-----------")
	var b1 B = cat
	b1.test2()
	fmt.Println("-----------")
	var c1 C = cat
	c1.test1()
	c1.test2()
	c1.test3()
}

// A1 接口
type A1 interface {
	test1()
}

// B 接口
type B interface {
	test2()
}

// C 接口
type C interface {
	A1
	B
	test3()
}

// Cat3 猫类型
type Cat3 struct { // 如果想要实现接口C,那么不止要实现接口C的方法,还要实现接口A1,B中的方法
}

func (c Cat) test1() {
	fmt.Println("test1()....")
}

func (c Cat) test2() {
	fmt.Println("test2()....")
}

func (c Cat) test3() {
	fmt.Println("test3()....")
}
