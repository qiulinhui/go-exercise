package structs

import "fmt"

// Interface2 空接口的实现
func Interface2() {
	/*
		空接口(interface{})
			不包含任何的方法,正因为如此,所有的类型都实现了空接口,因此空接口可以存储任意类型的数值.
	*/

	var a1 A2 = Cat2{"花猫"}
	var a2 A2 = Person1{"王二狗", 30}
	var a3 A2 = "哈哈"
	var a4 A2 = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	test1(a1)
	test1(a2)
	test1(a3)
	test1(a4)
	test1("Ruby")

	// map, key字符串,value任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "李小花"
	map1["age"] = 30
	map1["friend"] = Person1{"Jerry", 18}

	// 切片,存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 100, "abc")
}

// 接口A是空类型,理解为代表任意类型
func test1(a A2) {
	fmt.Println(a)
}

// 用空接口作为参数的典型
func test2(a interface{}) {

}

//A2 实现一个空接口
type A2 interface{}

// Cat2 结构体
type Cat2 struct {
	color string
}
