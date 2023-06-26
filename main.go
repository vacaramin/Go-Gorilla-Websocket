package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}
func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Websocket End Point")

}
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}

func main() {
	fmt.Println("Go Websocket")
}
