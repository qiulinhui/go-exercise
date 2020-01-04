package structs

import "fmt"

//Struct4 go语言中的OOP
func Struct4() {
	/*
		面向对象：OOP
		Go语言的结构体嵌套：
		1.模拟继承性：is -> a
		type A struct{
			field
		}

		type B struct{
			A //匿名字段
		}

		2.模拟聚合关系：has -> a

		type C struct{
			field
		}

		type D struct{
			c C //聚合关系
		}
	*/

	//1.创建父类的对象
	p1 := Person2{name: "张三", age: 30}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	//2.创建子类的对象
	s1 := Student4{Person2{"rose", 19}, "千锋教育"}
	fmt.Println(s1)

	s2 := Student4{Person2: Person2{name: "rose", age: 19}, school: "北京大学"}
	fmt.Println(s2)

	var s3 Student4
	s3.Person2.name = "王五"
	s3.Person2.age = 19
	s3.school = "清华大学"
	fmt.Println(s3)
}

//Person2 定义父类
type Person2 struct {
	name string
	age  int
}

//Student4 定义子类
type Student4 struct {
	Person2        //模拟继承结构
	school  string //子类的新增属性
}
