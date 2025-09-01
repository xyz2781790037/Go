package main
import "fmt"
import "time"
func s5h(){
	i := 2
	switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default :
		fmt.Println("It is after noon")
	}
	// 空接口可以接收 任何类型的值
	hpt := func(i interface{}){
		switch t := i.(type){
			case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
		}
	}
	hpt(1)
	// 类型开关 (type switch) 比较类型而非值。可以用来发现一个接口值的类型。 在这个例子中，变量 t 在每个分支中会有相应的类型。
}