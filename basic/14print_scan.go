package basic

import "fmt"
import "bufio"
import "os"

//PrintScan 标准输入和标准输出
func PrintScan() {
	/*
		输入和输出：
			fmt包：输入，输出

			输出：
				Print() //打印
				Printf() //格式化打印
				Println() //打印之后换行

			格式化打印占位符：
				%v,原样输出
				%T,打印类型
				%t,bool 类型
				%s,字符串
				%f,浮点
				%d,10 进制的整数
				%b,2 进制的整数
				%o,8 进制的整数
				%x,16 进制的整数 0-9 a-f
				%X,16 进制的整数 0-9 A-F
				%c,原样输出
				%p,原样输出
				...

			输入：
				Scanln()
				Scan()
			bufio包

	*/

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据", s1)
}
