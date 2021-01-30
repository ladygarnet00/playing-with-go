package utils

import "fmt"

type HTTPRequest struct {
	Method string
}

func LoopTilConditionWithBreak() {

	println("Loop til condition with break")

	// Loop til condition
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			println("break")
			break
		}
	}
}

func LoopTilConditionWithContinue() {

	println("Loop til condition with continue")

	// Loop til condition
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			continue
		}
		println("continuing")
	}
}

func ConditionalLoopsWithPostClauses() {

	println("Conditional loops with pause clauses")

	for i := 0; i < 5; i++ {
		println(i)
	}

	var j int
	for ; j < 5; j++ {
		println(j)
	}
}

func InfinitLoop() {

	println("Infinite loop")

	var i int
	for { // for ; ;
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}

func LoopOverSlices() {

	println("Loop over slices")

	//normal syntax
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for index, value := range slice {
		println(index, value)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}

	//print only keys
	for k := range wellKnownPorts {
		println(k)
	}

	//print only values, return keys but ignore them
	for _, v := range wellKnownPorts {
		println(v)
	}

}

func SwitchStatement() {
	r := HTTPRequest{Method: "GET"}
	switch r.Method {
	case "GET":
		println("Get request")
	case "PUT":
		println("Put request")
	case "POST":
		println("Post request")
	case "DELETE":
		println("Delete request")
	default:
		println("Unhandled method")
	}
}
