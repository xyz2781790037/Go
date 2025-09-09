package main
import "fmt"

type person struct {
    name string
    age  int
}
func newPerson(Name string) *person {
	p := person{name: Name}
	p.age = 42
	return &p
}
func structTest(){
	var per []person
	per = append(per,*newPerson("Bob"))
	per = append(per,person{name:"lfd",age:19})
	fmt.Println(per)
	fmt.Println(per[0].age)

	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
}