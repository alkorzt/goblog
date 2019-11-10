package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/footer.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	fmt.Println("Listening on port :5050")

	http.HandleFunc("/", indexHandle)

	http.ListenAndServe(":5050", nil)
}
