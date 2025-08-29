package main
// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

// 返回值的命名应当能反应其含义，它可以作为文档使用。

// 没有参数的 return 语句会直接返回已命名的返回值，也就是「裸」返回值。 
import "fmt"
// 在函数签名里写 (x, y int)、(res string, ok bool) 等，就是声明了命名返回值。它们在函数开始时已被声明并初始化为类型的零值。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func F(){
	fmt.Println(split(5))
}