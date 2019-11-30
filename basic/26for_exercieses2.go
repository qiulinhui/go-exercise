package basic

import "fmt"

// Exercieses2 九九乘法表
func Exercieses2() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
