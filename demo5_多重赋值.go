package main

import "fmt"

//go函数 可以返回多个值
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {

	a, b, c := 10, 20, 30

	//交换两个变量的值
	var tmp int
	tmp = a
	a = b
	b = tmp

	fmt.Printf("a = %d, b = %d,  c = %d\n", a, b, c)
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)

	i := 100
	j := 110
	//匿名变量，丢弃数据不处理
	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

	var d, e int
	c, d, e = test()
	fmt.Println("************", d, e)

}
