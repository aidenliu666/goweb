package main

import (
	"fmt"
	"net/http"
	"test/framework"
)

func main() {

	server := http.Server{
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}
	server.ListenAndServe()
	fmt.Println("abcdefd")
}
