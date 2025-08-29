package main
import "fmt"
import "math"

// 在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。
// pizza 和 pi 并未以大写字母开头，所以它们是未导出的。
// 在导入一个包时，你只能引用其中已导出的名字。 任何「未导出」的名字在该包外均无法访问。
// 执行代码，观察错误信息。
// 要修复错误，请将 math.pi 改名为 math.Pi，然后再试着执行一次。 


// 相当于是pravite和public的区别
func Add(x int,y int) int{
	return x + y;
}
// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。 
func addi(x,y int) int{
	return x + 2 * y;
}
func sub(x int ,y int) int{
	return x - y;
}
func ma() {
	// fmt.Println("hello world")
	a := 1
	b := 4
	// Sub(a,b)
	// rand.Intn(10)
	// math.Sqrt(67)
	// fmt.Printf("adhadaihhidhiahidahiahidahiadh\n")
	// fmt.Print("sawfsgf")

	fmt.Println(math.Pi)
	fmt.Printf("x + y = %d\n",Add(a,b))
	fmt.Printf("x + y = %d\n",addi(a,b))
	fmt.Println("x - y =",sub(a,b));
}