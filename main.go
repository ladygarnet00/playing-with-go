package main

import "fmt"

func main() {
	//Slices (builded over arrays)

	arr := [3]int{1, 2, 3}

	slice := arr[:]

	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4, 3, 56)
	fmt.Println(slice2)

	s2 := slice2[1:]
	s3 := slice2[:2]
	s4 := slice2[1:2]

	fmt.Println(s2, s3, s4)
}
