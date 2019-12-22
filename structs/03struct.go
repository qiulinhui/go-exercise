package structs

import "fmt"

// Struct2 匿名结构体
func Struct2() {
	/*
		匿名结构体和匿名字段：
		匿名结构体：没有名字的结构体，
			在创建匿名结构体时，同时创建对象

		匿名字段：一个结构体的字段没有字段名
		匿名函数：
	*/

	s1 := Student{name: "张三", age: 18}
	fmt.Println(s1.name, s1.age)
	s2 := struct {
		name string
		age  int
	}{
		name: "李四",
		age:  19,
	}
	fmt.Println(s2.name, s2.age)
	w2 := Worker{"李小花", 32}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)
}

// Worker 演示结构体匿名字段
type Worker struct {
	// name string
	// age  int
	string // 匿名字段
	int
	// string //匿名字段默认使用数据类型作为名字,那么匿名字段的类型就不能重复,否则会冲突
}

//Student 学生结构体
type Student struct {
	name string
	age  int
}
