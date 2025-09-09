// Go 支持匿名函数， 并能用其构造 闭包。 匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
package main
import "fmt"
// \intSeq 函数返回一个在其函数体内定义的匿名函数。 返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包。
func intSeq() func() int{
	i := 0
	return func() int{
		i++
		return i;
	}
}
// 我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值。
func Bb(){
	nextInt := intSeq()
	// 因为 Go 的闭包特性，这个 i 会一直保留在返回的函数中，不会随着 intSeq 执行完就销毁。
	fmt.Println(nextInt()) // i = 1
    fmt.Println(nextInt()) // i = 2
    fmt.Println(nextInt())
	newInts := intSeq()
    fmt.Println(newInts())
}