package main

import (
	"net/http"

	"github.com/Spirolina/todolist-restapi/handlers"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello world"))
	})
	http.HandleFunc("/add", handlers.CreateTodo)
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTodo(w, r)
		case http.MethodDelete:
			handlers.DeleteTodo(w, r)
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}

}
