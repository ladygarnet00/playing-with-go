package main

import (
	"fmt"
	"net/http"

	"github.com/ladygarnet00/playing-with-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	fmt.Println("Serving at port 3000")

	// // Loop til condition
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
	}

}
