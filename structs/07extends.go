package structs

import "fmt"

// Extends 继承
func Extends() {
	/*
		OOP中的继承：
		如果两个类(class)存在继承关系，其中一个是子类，另一个作为父类，那么：
			1.子类可以直接访问父类的属性和方法
			2.子类可以新增自己的属性和方法
			3.子类可以重写父类的方法（override,就是将父类已有的方法，重新实现)

		Go 语言的结构体嵌套：
			1.模拟继承性：is - a
				type A struct{
					field
				}
				type B struct{
					A //匿名字段
				}
			2.模拟聚合关系：has - a
				type C struct{
					field
				}

				type D struct{
					c C // 聚合关系
				}
	*/

	//1.创建 Person1 类型
	p1 := Person1{name: "王二狗", age: 30}
	fmt.Println(p1.name, p1.age) //父类对象 访问父类的字段属性
	p1.eat()                     //父类对象 访问父类的方法

	//2.创建 Student1 类型
	s1 := Student1{
		Person1{name: "Ruby", age: 18}, "千锋教育",
	}
	fmt.Println(s1.name)   //s1.Person.name
	fmt.Println(s1.age)    // 子类对象可以直接访问父类的字段，（其实就是提升字段）
	fmt.Println(s1.school) //子类对象访问自己新增的字段
	s1.eat()               // 子类对象访问父类的方法
	s1.study()             //子类对象访问自己新增的方法
	s1.eat()               //子类对象访问重写的方法
}

// Person1 1.定义一个“父类”
type Person1 struct {
	name string
	age  int
}

// Student1 2.定义一个“子类”
type Student1 struct {
	Person1 //结构体嵌套模拟继承性
	school  string
}

//3.方法
func (p Person1) eat() {
	fmt.Println("父类的方法，吃窝窝头")
}

func (s Student1) study() {
	fmt.Println("子类新增的方法，学生学习啦。。。")
}

func (s Student1) eat() {
	fmt.Println("子类重写的方法，吃炸鸡喝啤酒。。。")
}
