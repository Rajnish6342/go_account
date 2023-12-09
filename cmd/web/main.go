package main

import (
	"fmt"
	"net/http"

	"github.com/Rajnish6342/go_app/pkg/handlers"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Server Started on %s 🚧", port)
	_ = http.ListenAndServe(port, nil)
}
