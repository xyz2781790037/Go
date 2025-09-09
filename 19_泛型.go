package main
import "fmt"
func MapKeys[K comparable,V any](m map[K]V) []K  {
	r := make([]K,0,len(m))
	for k := range m{
		r = append(r, k)
	}
	return r
}
type List[T any]struct{
	head,tail *element[T]
}
type element[T any]struct{
	next *element[T]
	val T
}
func (lst *List[T]) Push(v T){
	if(lst.tail == nil){
		lst.tail = &element[T]{val: v}
		lst.head = lst.tail
	}else{
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T{
	var elemt []T
	for phead := lst.head;phead != nil;phead = phead.next{
		elemt = append(elemt, phead.val)
	}
	return elemt
}
func ListTest(){
	var m = map[int]string{1:"2",2:"4",3:"6"}
	fmt.Println("keys m:", MapKeys(m))
	// 只调用函数而忽略返回值，或者 让编译器不报未使用的函数/变量错误。
	// _ = MapKeys[int, string](m)// 告诉编译器：我故意丢掉
	list := List[int]{}
	list.Push(12)
	list.Push(34)
	list.Push(56)
	list.Push(78)
	fmt.Println("list:",list.GetAll())
}