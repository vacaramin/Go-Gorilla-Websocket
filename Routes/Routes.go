package Routes

import "net/http"

func SetupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}
