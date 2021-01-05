package main

import "fmt"

func main() {

	//Structs

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	//long initialization
	var u user
	u.ID = 8
	u.FirstName = "Carolina"
	u.LastName = "Martinez"
	fmt.Println(u)
	fmt.Println(u.ID)

	//short initialization
	u2 := user{ID: 1,
		FirstName: "Lechuga",
		LastName:  "Martinez",
	}
	fmt.Println(u2)

}
