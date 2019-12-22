package function

import "fmt"

// Type 函数的类型
func Type() {
	/*
		函数的类型因函数本身的参数和返回值类型不同而有所不同
	*/
	fmt.Printf("%T\n", func10)
	fmt.Printf("%T\n", func11)
	fmt.Printf("%T\n", func12)
	fmt.Printf("%T\n", func13)
}

func func10() {

}
func func11(a int) int {
	return 0
}

func func12(a float64, b, c int) (int, int) {
	return 0, 0
}
func func13(a, b string, c, d int) (string, int, float64) {
	return "", 0, 0
}
