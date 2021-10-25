package main

import "fmt"

// type Iname interface {
// 	Mname()
// }
// type St struct {}
// func (St) Mname() {}
// func main() {
// 	var t *St
// 	if t == nil {
// 		fmt.Println("t is nil")
// 	} else {
// 		fmt.Println("t is not nil")
// 	}
// 	fmt.Printf("\n")
// 	var i Iname = t
// 	fmt.Printf("%T\n", i)
// 	fmt.Printf("%v\n", i)
// 	fmt.Println(i, nil, i == nil) // 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil
// 	if i == nil {
// 		fmt.Println("i is nil")
// 	} else {
// 		fmt.Println("i is not nil")
// 	}
// 	fmt.Printf("i is nil pointer:%v",i == (*St)(nil))
// }

type Shape interface {
	Area() float32
}

type Rect struct {
	width  float32
	height float32
}

func (r Rect) Area() float32 {
	return r.width * r.height
}

func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectange s", s.Area())
	fmt.Println("s == r is", s == r)
}
