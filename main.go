package main

import "fmt"

// 1.变量
// 2.函数
// 3.循环
// 4.外部调用


func main() {
	fmt.Println("gggfff")
	var a int = 10	// 声明定义为int
	var b = 11		// 声明自动判定变量数据类型
	c := 12			// 未声明，自动判定变量数据类型
	fmt.Println(a, b, c)

	// 多变量初始化
	var e, f, g = 21, 22, 23
	fmt.Println(e, f, g)
	h, i, j := 31, 32, 33
	fmt.Println(h, i, j)

	// 条件判断
	if 1 < f {
		fmt.Println("e < f")
	}

	// for循环
	for m := 0; m < 10; m++ {
		fmt.Println(m)
	}

}
