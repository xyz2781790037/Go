package main
import "fmt"
func addb(a int,b int){
	fmt.Println("result is",a + b)
}
// 函数可以返回任意数量的返回值。

// swap 函数返回了两个字符串。 
func Sub(a, b int) (int,int,int){
	addb(a,b)
	c := 10
	return c - (a - b),c,b;
}
func x4(){
	ma()
	a := 2
	b := 3
	var d,e,f int
	d,e,f = Sub(a,b)
	fmt.Println(d,e,f)
	F()
	BBB()
}