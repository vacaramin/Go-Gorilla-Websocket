package main

import (
	"fmt"
	"net/http"

	"github.com/vacaramin/Go-Gorilla-Websocket/Routes"
)

func main() {
	fmt.Println("Go Websocket")
	Routes.SetupRoutes()

	http.ListenAndServe(":8080", nil)

}
