package error

import "fmt"

//Panic 恐慌
func Panic() {
	/*
		panic:词义 “恐慌”
		recover：“恢复”
		go 语言利用 panic(), revocer()实现程序中的极特殊的异常的处理
			panic(),让当前程序进入恐慌，中断程序的运行
			recover(),让程序恢复，必须在 defer 函数中执行
	*/
	funA()
	defer myprint("defer main:3...")
	funB()
	defer myprint("defer main:4...")
	fmt.Println("main over....")
}

func myprint(s string) {
	fmt.Println(s)
}

func funA() {
	fmt.Println("我是一个函数 funcA()...")
}

func funB() { //外围函数
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复了")
		}
	}()
	fmt.Println("我是函数funB()...")
	defer myprint("defer funB():1......")
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//让程序中断
			panic("funB函数恐慌了\n")
		}
	} //当外围函数的代码中发生了运行恐慌，只有其中所有的已经defer的函数全部都执行完毕后，该运行恐慌才会真正被扩展至调用处。
	defer myprint("defer funB():2.....")
}
