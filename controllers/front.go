package controllers

import (
	"net/http"
)

//router for users
func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}
