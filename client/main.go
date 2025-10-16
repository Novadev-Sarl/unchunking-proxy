package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Started chunked client to port 8080...")

	http.Post("http://localhost:8080/test", "application/json", nil)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
