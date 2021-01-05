package main

import (
	"fmt"

	"github.com/ladygarnet00/playing-with-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Carolina",
		LastName:  "Martinez",
	}

	fmt.Println(u)
}
