package main

import (
	"fmt"
	"net/http"

	"go.mod/api"
)

func main() {
	srv := api.NewServer()

	fmt.Println("The server is running on port :8000")
	http.ListenAndServe(":8000", srv)
}
