package basic

import "fmt"

// Iota2 Iota演示2
func Iota2() {
	const (
		A = iota   // 0
		B          // 1
		C          // 2
		D = "haha" // haha iota = 3
		E          //haha iota = 4
		F = 100    // 100 iota = 5
		G          // 100 iota = 6
		H = iota   // 7
		I          // 8
	)
	const (
		J = iota
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)

}
