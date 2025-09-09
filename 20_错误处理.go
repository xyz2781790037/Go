package main
// 符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。 这与 Java、Ruby 使用的异常（exception） 以及在 C 语言中有时用到的重载 (overloaded) 的单返回/错误值有着明显的不同。 Go 语言的处理方式能清楚的知道哪个函数返回了错误，并使用跟其他（无异常处理的）语言类似的方式来处理错误。
import (
    "errors"
    "fmt"
)
func f1(arg int)(int,error){
	if arg == 42{
		return -1,errors.New("can't work with 42")
	}
	return arg + 3,nil
}
type argError struct {
	arg int
	prob string
}
func (e *argError)Error() string{
	return fmt.Sprintf("%d - %s",e.arg,e.prob)

}
func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func errorTest(){
	for _,i := range []int{7,42}{
		if r,e := f1(i);e != nil{
			fmt.Println("f1 failed:",e)
		}else{
			fmt.Println("f1 worked:",r)
		}
	}
	for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}