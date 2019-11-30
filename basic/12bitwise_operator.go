package basic

import "fmt"

func bitwise() {
	/*
		位运算符：
			将数值，转为二进制后，按位操作
		按位&：
			对应位的值如果都为 1 才为 1，有一个为 0 就为 0
		按位|:
			对应位的值如果都为 0 才为 0，有一个为 1 就为 1
		异或^:
			二元：a^b
				对应的值不同为 1，相同为 0
			一元：^a
				按位取反：1 ---> 0, 0 ---> 1
		位清空：&^
			对于 a&^ b
				对于 b 上的每个数值
				如果为 0，则取 a 对应位上的数值
				如果为 1，则结果位就取 0
		位移运算符：
		<<: 按位左移，将 a 转为二进制，向左移动 b 位
			a << b
		>>:按位右移，将 a 转为二进制，向右移动 b 位
			a >> b
	*/

	a := 60
	b := 13
	/*
		a := 60 0011 1100
		b := 13 0000 1101
		&       0000 1100
		|       0011 1101
		^       0011 0001
		&^      0011 0000
	*/

	fmt.Printf("a:%d, %b\n", a, a)
	fmt.Printf("b:%d, %b\n", b, b)

	res1 := a & b
	fmt.Println(res1) // 12

	res2 := a | b
	fmt.Println(res2) // 61

	res3 := a ^ b
	fmt.Println(res3) // 49

	res4 := a &^ b
	fmt.Println(res4) // 48

	res5 := ^a
	fmt.Println(res5) // -61

	c := 8
	res6 := c << 2
	fmt.Println(res6) //32

	res7 := c >> 2
	fmt.Println(res7) //2
}
