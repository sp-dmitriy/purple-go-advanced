package main

import (
	"fmt"
	"net/http"
)


func main() {
	router := http.NewServeMux()
	NewRandHandler(router)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server started on port 8081")
	if err := server.ListenAndServe()
	err != nil {
		fmt.Println("Error starting server:", err)
	}
}
