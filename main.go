package main

import (
	"fmt"
	"goweb/framework"
	"goweb/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	server.ListenAndServe()
	fmt.Println("abcdeffd")
}
