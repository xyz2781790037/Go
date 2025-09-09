// 协程(goroutine) 是轻量级的执行线程。
package main

import (
    "fmt"
	"sync"
)
func f(from string) {
	for i := 0;i < 3;i++{
		fmt.Println(from,":",i)
	}
}
func Coroutine(){
	f("direct")
	// 使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
	var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        f("goroutine")
    }()

    go func(msg string) {
        defer wg.Done()
        for i := 0;i < 3;i++{
		fmt.Println(msg,":",i)
	}
    }("going")

	// go f("goroutine")
	// go func(msg string){
	// 	fmt.Println(msg)
	// }("going")
	wg.Wait()
	fmt.Println("done")
}