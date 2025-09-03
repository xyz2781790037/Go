package main
import "fmt"
func mapTest(){
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 5
	fmt.Println("map:", m)
	v1 := m["k2"]
	fmt.Println(v1)

	fmt.Println("len:", len(m))
	// 内建函数 delete 可以从一个 map 中移除键值对。
	delete(m, "k2")
	fmt.Println("map:", m)
	// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略。
	_, prs := m["k2"]
    fmt.Println("prs:", prs)
	// 你也可以通过下边的语法在一行代码中声明并初始化一个新的 map。
	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}