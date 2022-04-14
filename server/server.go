package server

import (
	"fmt"
	"github.com/KyriakosMilad/go-rest-cms/user"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})
	user.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
