package structs

import "fmt"

// Method 方法演示
func Method() {
	/*
		一个方法就是包含了接收者的函数，接收者可以是命名类型或者结构体类型的一个值或一个指针。
		所有给定类型的方法属于该类型的方法集

		语法：
			func(接收者) 方法名 (参数列表)(返回值列表) {

			}
		总结：method 同函数类似，区别是要有接收者。（也就是调用者）
		对比函数：
			A：意义
				方法：某一个类别的行为功能，需要指定的接收者调用
				函数：一段独立功能的代码，可以直接调用
			B：语法
				方法：方法名可以相同，只要接收者不同
				函数：命名不能冲突
	*/
	w1 := Worker2{name: "王二狗", age: 30, sex: "男"}
	w2 := &Worker2{name: "Ruby", age: 20, sex: "女"}
	w1.work()
	w2.work()
	w2.rest()
	w1.rest()
	w2.printInfo()
}

// Worker2 定义一个工人结构体
type Worker2 struct {
	//字段
	name string
	age  int
	sex  string
}

// Cat 猫咪
type Cat struct {
	color string
	age   int
}

//定义行为方法
func (w Worker2) work() {
	fmt.Println(w.name, "在工作...")
}

//休息
func (w *Worker2) rest() {
	fmt.Println(w.name, "在休息")
}

func (w *Worker2) printInfo() {
	fmt.Printf("工人姓名：%s，工人年龄：%d，工人性别:%s\n", w.name, w.age, w.sex)

}

func (p *Cat) printInfo() {
	fmt.Printf("猫咪的颜色：%s，年龄：%d\n", p.color, p.age)
}
