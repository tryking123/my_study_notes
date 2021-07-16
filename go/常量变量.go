package main

import "fmt"

/*
Go 是静态类型语⾔，不能在运⾏期改变变量类型。
使⽤关键字 var 定义变量，⾃动初始化为零值。如果提供初始化值，可省略变量类型，由
编译器⾃动推断
*/
var x int
var f float32 = 1.6

var s, n = "abc", 123
var a1 uint = 5

func main() {
	//x := 123 报错 ==> 定义了的变量一定要用,不然报错
	//在函数内部，可⽤更简略的 ":=" ⽅式定义变量。
	x := x + 123
	fmt.Println(x)
	fmt.Println(f)
	fmt.Println(s, n, a1)
	fmt.Println("===============")
	data, i := [3]int{0, 1, 2}, 0
	fmt.Println(i)
	fmt.Println(data[0])
	i, data[i] = 2, 100 // (i = 0) -> (i = 2), (data[0] = 100)
	fmt.Println(i)
	fmt.Println(data[0])
	fmt.Println("===============")

	_, s := test()
	fmt.Println(s)
	fmt.Println("===============")
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	fmt.Println("===============")

}
func test() (int, string) {
	return 1, "abc"
}
