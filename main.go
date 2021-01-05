package main

import (
	"fmt"
)

func main() {

	//pointer
	var firstName *string = new(string)

	*firstName = "carolina"

	fmt.Println(*firstName)

	secondName := "Carolina"
	fmt.Println(secondName)

	ptr := &secondName
	fmt.Println(ptr, *ptr)

	secondName = "Lechuga"
	fmt.Println(ptr, *ptr)
}
