package basic

import (
	"fmt"
	"math/rand"
	"time"
)

// Random 生成随机数
func Random() {
	/*
		生成随机数 random:
			伪随机数，根据一定的算法公式算出来的
			math/rand
	*/

	// step1: 设置种子数，可以设置为时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//step2: 调用生成随机数的函数
		fmt.Println("--->", rand.Intn(100))
	}
}
