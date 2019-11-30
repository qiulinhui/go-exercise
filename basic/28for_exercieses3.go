package basic

import (
	"fmt"
	"math"
)

// Exercieses3 求水仙花数
func Exercieses3() {
	/*
		水仙花数：三位数[100,999]
		每个位上的数字的立方和，刚好等于该数字本身。那么就叫水仙花数，4个
		比如：153
			1*1*1 + 5*5*5 + 3*3*3 = 1 + 125 + 27 = 153
	*/
	for i := 100; i < 1000; i++ {
		x := i / 100     //百位
		y := i / 10 % 10 //十位
		z := i % 10      //个位

		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
}
