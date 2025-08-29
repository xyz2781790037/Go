// := 语法是声明并初始化变量的简写， 例如 var f string = "short" 可以简写为右边这样。
package main

import "fmt"

func BBB() {

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