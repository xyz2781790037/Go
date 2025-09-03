package main
import "fmt"
func arrTest(){
	var a [5]int
	fmt.Println("a:",a)
	a[3] = 20
	fmt.Println("a:",a)
	fmt.Println("a[5]:",a[4])
	fmt.Println("len:",len(a))
	b := [5]int{1,2,3,4,5}
	fmt.Println("b:",b)
	var c [3][3]int
	for i:= 0;i < 3;i++{
		for j := 0;j < 3;j++{
			c[i][j] = i + j
		}
	}
	fmt.Println("c:",c)
}