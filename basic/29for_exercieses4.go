package basic

import (
	"fmt"
	"math"
)

// Exercieses4 求水仙花数
func Exercieses4() {
	/*
		打印 2-100 内的素数
	*/

	for i := 2; i <= 100; i++ {
		flag := true //记录 i 是否是素数
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				flag = false //不是素数
				break
			}
		}

		if flag {
			fmt.Println(i)
		}
	}
}
