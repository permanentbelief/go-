//package main

// import (
// 	"fmt"
// 	"os" // 包含一些函数和变量 以及与平台无关的方式和操作系统打交道
// 	// os.Args是一个字符串slice  动态容量的顺序数组
// )

// func main() {
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

package main

import "fmt" //fmt.Println()

func main() {
	fmt.Println("hello world")
}
