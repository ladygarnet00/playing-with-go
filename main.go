package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello from module, Gophers")

	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	//implicit initialization
	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	//(3+4i)
	c := complex(3, 4)
	fmt.Println(c)

	//3 4
	r, im := real(c), imag(c)
	fmt.Println(r, im)

}
