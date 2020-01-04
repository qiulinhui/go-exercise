package error

import (
	"fmt"
	"net"
)

// Error3 error 演示
func Error3() {
	addr, err := net.LookupHost("www.jintian.com")
	fmt.Println(err)
	fmt.Println(addr)
	if ins, ok := err.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("操作超时")
		} else if ins.Temporary() {
			fmt.Println("临时性错误")
		} else {
			fmt.Println("通常错误")
		}
	}
}
