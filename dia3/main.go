package main

import (
	"dia3/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handler.GetUsersHandler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
