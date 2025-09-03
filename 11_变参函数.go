// 可变参数函数。 在调用时可以传递任意数量的参数。 例如，fmt.Println 就是一个常见的变参函数。
package main

	

import "fmt"

// 这个函数接受任意数量的 int 作为参数。

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)

}
func totalSum(){
	sum(1, 2)
    sum(1, 2, 3)
	nums := []int{1,2,3,4}
	sum(nums...)
}