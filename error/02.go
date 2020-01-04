package error

import (
	"errors"
	"fmt"
)

// Error2 error 演示
func Error2() {
	/*
		error:内置的数据类型,内置的接口
		定义方法:Error() string
		使用 go 语言提供好的包:
			errors 包下的函数:New(),创建一个 error 对象
			fmt包下的Errorf()函数：
				func Errorf(format string, a ...interface{}) error
	*/
	// 1. 创建一个 error 数据
	err1 := errors.New("自定义的错误")
	fmt.Println(err1)
	fmt.Printf("%T", err1) //*errors.errorString

	//2. 创建另一个 error 的方法
	err2 := fmt.Errorf("错误的信息码：%d", 100)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)
}

//设计一个函数：验证年龄是否合法，如果为负数，就返回一个 error
func checkAge(age int) error {
	if age <= 0 {
		//返回 error 对象
		return errors.New("年龄不合法")
	}
	fmt.Println("年龄是：", age)
	return nil
}
