package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
