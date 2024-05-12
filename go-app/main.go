package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Golang APP is Running")
	})
	fmt.Println("Golang APP is Running...")
	http.ListenAndServe(":8080", nil)
}
