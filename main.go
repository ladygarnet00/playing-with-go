package main

import (
	"net/http"

	"github.com/ladygarnet00/playing-with-go/controllers"
	"github.com/ladygarnet00/playing-with-go/models"
	"github.com/ladygarnet00/playing-with-go/utils"
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

	models.GetUsers()

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

	utils.SwitchStatement()

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
