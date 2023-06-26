package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Go Websocket")
}
