package main

import (
	"fmt"

	"github.com/ladygarnet00/playing-with-go/models"
)

func main() {

	u1 := models.User{
		ID:        1,
		FirstName: "Carolina",
		LastName:  "Martinez",
	}

	u2 := models.User{
		ID:        2,
		FirstName: "Juan",
		LastName:  "Diaz",
	}

	if u1.ID == u2.ID {
		println("Same ID!")
	} else if u1.FirstName == u2.FirstName {
		println("Same first name")
	} else if u1.LastName == u2.LastName {
		println("Same last name")
	} else {
		println("Different users")
	}

	if u1.ID != u2.ID {
		println("Not same ID!")
	} else {
		println("Same Users")
	}

	println(u1.ID)
	println(u2.ID)

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
