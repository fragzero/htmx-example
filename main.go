package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

var port = ":8080"

type TodoList struct {
	Title       string
	Description string
}

func main() {
	fmt.Printf("Starting server on port %s\n", port)

	homepage := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))

		todoList := map[string][]TodoList{
			"TodoList": {
				{Title: "1", Description: "Fazer compras"},
				{Title: "2", Description: "Pagar conta de luz"},
				{Title: "3", Description: "Tirar férias"},
			},
		}
		template.Execute(w, todoList)
	}

	addTodo := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "todo-list-element", TodoList{Title: title, Description: description})
	}

	http.HandleFunc("/", homepage)
	http.HandleFunc("/add-todo", addTodo)

	log.Fatal(http.ListenAndServe(port, nil))
}
