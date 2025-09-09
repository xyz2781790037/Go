package main
// Go支持对于结构体(struct)和接口(interfaces)的 嵌入(embedding) 以表达一种更加无缝的 组合(composition) 类型
import "fmt"

type base struct {
    num int
}
func (b base) describe() string{
	return fmt.Sprintf("base with num=%v",b.num)
}
type container struct{
	base
	str string
}
func embeddingTest(){
	co := container{
		base{
			1},
		"some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println(co)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:",co.describe())
	fmt.Println("describe:",co.str)
	// fmt.Scanln()
	type des interface{
		describe() string
	}
	var d des = co
	fmt.Println("describer:",d.describe())
}