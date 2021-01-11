package main

import (
	"errors"
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

	//Functions, arguments, parameters, returning data

	startWebServer()

	port := 3000
	retries := 3
	startWebServerWithPortRetries(port, retries)

	isStarted := startWebServerReturn(port)
	fmt.Println("Is server started?", isStarted)

	err := startWebServerOnError(port)
	fmt.Println(err)

	port, err2 := startWebServerMultipleReturn(port)
	fmt.Println(err2, port)

	//if i dont need the first return value. Ignore the first result
	_, err3 := startWebServerMultipleReturn(port)
	fmt.Println(err3)

}

func startWebServer() {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started")
}

func startWebServerWithPortRetries(port, numberRetries int) {
	fmt.Println("Starting server at port", port, numberRetries)
	//do important things
	fmt.Println("Server started")
}

func startWebServerReturn(port int) bool {
	fmt.Println("Starting server at port", port)
	//do important things
	fmt.Println("Server started")

	return true
}

func startWebServerMultipleReturn(port int) (int, error) {
	fmt.Println("An error has ocurred", port)
	//do important things
	fmt.Println("Server shut")

	return port, nil
}

func startWebServerOnError(port int) error {
	fmt.Println("An error has ocurred", port)
	//do important things
	fmt.Println("Server shut")

	return errors.New("An error has ocurred")
}
