package error

import "os"

import "fmt"

// Error1 error 演示
func Error1() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.Op:", ins.Op)
			fmt.Println("2.Path：", ins.Path)
			fmt.Println("3.Err：", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功...")
}
