package main

import "fmt" // 导入包后必须使用

func main() {
	//变量 程序运行期间可以改变的量

	//1.声明格式 var 变量名 变量类型
	//2.只是声明没有初始化的变量 默认值为0
	//3.同一个{}.声明的变量名是唯一的

	var a int
	fmt.Println("aaa")
	fmt.Println("a = ", a)

	//var b, c int
	a = 10

	var d int = 100
	fmt.Println(d)

	//自动推导类型

	c := 30
	fmt.Printf("c type is %T\n", c)

}
