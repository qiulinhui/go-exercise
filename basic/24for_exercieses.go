package basic

import "fmt"

//ForExercieses for 练习题
func ForExercieses() {
	/*
		练习1：打印 58-23 数字
		练习2：+1+2+3+4...+100
		练习3：打印 1-100 内能够被 3 整除，但是不能被 5 整除的数字，统计被打印的数字的个数，每行打印 5 个
	*/

	for i := 58; i >= 23; i-- {
		fmt.Println(i)
	}

	fmt.Println("--------------------")

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100的和：", sum)

	fmt.Println("--------------------")

	count := 0

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, "\t")
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Println("count ->", count)
}
