package error

import (
	"fmt"
	"path/filepath"
)

// Error4 error 演示
func Error4() {
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err) //syntax error in pattern
		return
	}
	fmt.Println("files:", files)
}
