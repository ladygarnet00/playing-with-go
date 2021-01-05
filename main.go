package main

import "fmt"

func main() {
	//Arrays

	//verbose way
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println(array)
	fmt.Println(array[1])

	//implicit initialization
	array2 := [3]int{1, 2, 3}
	fmt.Println(array2)
}
