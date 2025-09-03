package main
import "fmt"
func Range(){
	nums := []int{2,3,4}
	sum := 0
	for _,num := range nums{
		sum += num
	}
	fmt.Println("sum:",sum)
	for i,num := range nums{
		if(num == 3){
			fmt.Println("index:",i)
		}
	}
	kvs := map[string]string{"a":"ba","abc":"cba"}
	for k, v := range kvs {
        fmt.Println(k,"->", v)
    }
	for k := range kvs {
        fmt.Println("key:", k)
    }
	for i, c := range "go" {
        fmt.Println(i, c)
    }
}