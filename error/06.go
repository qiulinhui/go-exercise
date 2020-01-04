package error

import "fmt"

// Error6 error 演示
func Error6() {
	length, width := -6.7, 9.1
	area, err := rectArea(length, width)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError2); ok {
			if err.lengthNegative() {
				fmt.Printf("长度，%.2f,小于零值\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("宽度，%.2f,小于零值\n", err.width)
			}
		}
		return
	}

	fmt.Println("矩形的面积是", area)
}

//1.定义一个结构体，表示错误的类型
type areaError2 struct {
	msg    string  //错误的描述
	length float64 //发生错误时矩形的长度
	width  float64 //发生错误时矩形的宽度
}

func (e *areaError2) Error() string {
	return e.msg
}

func (e *areaError2) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError2) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	msg := ""

	if length < 0 {
		msg = "长度小于零"
	}
	if width < 0 {
		if msg == "" {
			msg = "宽度小于零"
		} else {
			msg += "宽度也小于零"
		}
	}
	if msg != "" {
		return 0, &areaError2{msg, length, width}
	}
	return length * width, nil
}
