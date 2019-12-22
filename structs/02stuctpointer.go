package structs

import "fmt"

//Sp 结构体指针
func Sp() {

	//1.结构体是值类型的数据
	p1 := Person{"王二狗", 30, "男", "北京市"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)

	p2.name = "李小花"
	fmt.Println(p2)
	fmt.Println(p1)

	//2.定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1)
	fmt.Println(*pp1)

	//(*pp1).name = "王五"
	pp1.name = "王五"
	fmt.Println(pp1)
	fmt.Println(p1)

	//使用内置函数new(),go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)
	fmt.Println(pp2)
	//(*pp2).name
	pp2.name = "Jerry"
	pp2.age = 20
	pp2.sex = "男"
	pp2.address = "上海市"
	fmt.Println(pp2)

	//new()，不是nil，空指针，指向了新分配的类型的内存空间，里面存储的零值。
}
