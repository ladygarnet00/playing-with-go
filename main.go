package main

import (
	"net/http"

	"github.com/ladygarnet00/playing-with-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
