package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on port :5050")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>hello...</h1>")
	})

	http.ListenAndServe(":5050", nil)
}
