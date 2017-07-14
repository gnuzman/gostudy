package main

import "fmt"

// 1.变量
// 2.函数
// 3.循环
// 4.外部调用
func test1()(string) {
	return "test"
}

func test2(p1 string, p2 string)(string, string) {
	return p2, p1
}

func test3(p1 string, p2 int)(int, string) {
	return p2, p1
}

// 结构体
type Books struct {
	title string
	author string
	subject string
	book_id int
}

// 结构体作为函数参数和返回值
func testStruct(book *Books) *Books {
	book.author = "zzh"
	book.subject = "you"
	book.title = "nice"
	book.book_id = 1

	return book
}

func main() {
	fmt.Println("ggg")
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

	fmt.Println(test1())
	fmt.Println(test2("1", "2"))
	fmt.Println(test3("10", 11))


	// 函数指针作为参数和返回值
	var myBook Books
	var myPtr *Books = testStruct(&myBook)
	fmt.Println(myBook.book_id, myBook.author, myBook.title, myBook.subject)
	fmt.Println(myPtr.book_id, myPtr.author, myPtr.title, myPtr.subject)

}
