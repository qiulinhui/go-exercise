package structs

import "fmt"

// Interface 接口的实现
func Interface() {
	/*
		接口：interface
			在 Go 中，接口是一组方法签名。
			当某个类型为这个接口中的所有方法提供了方法的实现，它被称为实现接口
			Go 语言中，接口和类型的实现关系，是非侵入式

		1.当需要接口类型的对象时,可以使用任意实现类对象替代
		2.接口对象不能访问实现类中的属性

		多态:一个事物的多种形态
			go语言通过接口模拟多态
			就一个接口的实现
				1.看成实现本身的类型,能够访问实现类中的属性和方法
				2.看成是对应的接口类型,那就只能够访问接口中的方法
		接口的用法:
			1.一个函数如果接受接口类型作为参数,那么实际上可以传入该接口的任意实现类型对象
	*/

	//1.创建 Mouse
	m1 := Mouse{"罗技小红"}
	//2.创建 FlashDisk
	f1 := FlashDisk{"闪迪64G"}
	fmt.Println(m1.name)
	testInterface(m1)
	testInterface(f1)

	var usb USB
	usb = m1
	usb.start()
	usb.end()
	// fmt.Println(usb.name) //不能访问
}

//USB 1.定义接口
type USB interface {
	start() // USB设备开始工作
	end()   //USB设备结束工作
}

// Mouse 2.实现类
type Mouse struct {
	name string
}

// FlashDisk 实现类
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标，准备就绪，可以开始工作了，点点点")
}

func (m Mouse) end() {
	fmt.Println(m.name, "结束工作，可以安全退出了。。。。")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "准备开始工作，可以进行数据的存储。。")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "可以弹出。。")
}

//3.测试方法
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
