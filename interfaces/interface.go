package main

import "fmt"

type I interface {
	Show()
}

type F struct {
	f float64
}

type S struct {
	S string
}

func (x F) Show() {
	fmt.Println(x.f)
}

func (x *S) Show() {
	fmt.Println(x.S)
}

func Ishow(i I) {
	i.Show()
}

func typecheck(i interface{}) {
	switch i.(type) {
	case F:
		fmt.Println("Yay, type F!")
	case S:
		fmt.Println("Yay, type S!")
	case *S:
		fmt.Println("Yay, type pointer to S!")
	default:
		fmt.Println("Unrecognized type")
	}
}

func main() {
	var i I

	i = F{0.35}
	i.Show()
	Ishow(i) // The interface is the argument
	typecheck(i)

	i = &S{"test"}
	i.Show()
	Ishow(i) // The interface is the argument
	typecheck(i)
}
