package basic

import "fmt"

// Continue for循环中的 continue
func Continue() {
	/*
		循环结束：
			循环条件不满足，循环自动结束
			但是可以通过 break 和 continue 来强制结束循环
		循环控制语句：
		break：彻底的结束循环。。。终止
		continue：结束了某一次循环，下次继续。。。中止

		注意点：多层循环嵌套，break和continue,默认结束的是里层循环
		如果想结束指定的某个循环，可以给循环贴标签（起名）
		break 循环标签名
		continue 循环标签名
	*/

	for i := 1; i <= 10; i++ {
		if i == 5 {
			// break
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("---------------------")

out:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			break out
			// continue out
		}
	}
}
