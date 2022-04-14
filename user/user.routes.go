package user

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/users/register", registerHandler)
}
