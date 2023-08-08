package Routes

import "net/http"

// All routes will be set here
func SetupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndPoint)
}
