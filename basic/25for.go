package basic

import "fmt"

// For3 多层 for 循环
func For3() {
	for i := 1; i <= 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
