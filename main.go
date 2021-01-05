package main

import "fmt"

const pi = 3.14

const (
	first  = 1
	second = "A"
	third  = true
)

const (
	a = iota + 4
	b = 2 << iota
	c = iota
	d = iota
)

const (
	f = iota + 4
	g
)

const (
	h = iota
)

func main() {

	//Using Iota and Constant expressions
	fmt.Println(pi, first, second, third)

	fmt.Println(a, b, c, d)
	fmt.Println(f, g)
	fmt.Println(h)
}
