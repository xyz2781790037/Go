package main

import "fmt"

func cycle() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
	for i := 0;i <= 10;i++{
		fmt.Print("11")
		// 在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用。
		// Go 没有三目运算符， 即使是基本的条件判断，依然需要使用完整的 if 语句。
		if j:= 5;i > j{
			fmt.Print("\n")
		}
	}
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}