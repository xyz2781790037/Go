// := 语法是声明并初始化变量的简写， 例如 var f string = "short" 可以简写为右边这样。
package main

import "fmt"
import "math"
// const 用于声明一个常量。

// const 语句可以出现在任何 var 语句可以出现的地方
	
// 常数表达式可以执行任意精度的运算

// 数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化。
func BBB() {

	const x = 100.222
	fmt.Println(x * 2);
	const y = 10.00034
	fmt.Printf("%g\n",math.Ceil(y))
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "short"
    fmt.Println(f)
}