package main

import (
	"fmt"
)

func main() {

	loopTilConditionWithBreak()
	loopTilConditionWithContinue()
	conditionalLoopsWithPostClauses()
	infinitLoop()
	loopOverSlices()

	//controllers.RegisterControllers()
	//http.ListenAndServe(":3000", nil)
	//fmt.Println("Serving at port 3000")

}

func loopTilConditionWithBreak() {

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

func loopTilConditionWithContinue() {

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

func conditionalLoopsWithPostClauses() {

	println("Conditional loops with pause clauses")

	for i := 0; i < 5; i++ {
		println(i)
	}

	var j int
	for ; j < 5; j++ {
		println(j)
	}
}

func infinitLoop() {

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

func loopOverSlices() {

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
