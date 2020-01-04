package structs

import "fmt"

//Type2 type关键字演示
func Type2() {
	var s Student5
	s.Person.name = "王二狗"
	s.Person.show()
	fmt.Printf("%T,%T\n", s.Person, s.People)
	s.People.name = "李小花"
	s.People.show()
	s.People.show2()
	s.Person.show()
}

//Person3 类型
type Person3 struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person----->", p.name)
}

// func (p People) show() {} //报错：重复定义了show方法
func (p People) show2() {
	fmt.Println("People------>", p.name)
}

// People 类型别名
type People = Person

// Student5 学生类型
type Student5 struct {
	Person
	People
}
