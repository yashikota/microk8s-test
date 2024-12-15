package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r* http.Request) {
	fmt.Fprintf(w, "Hello, World! v2")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
