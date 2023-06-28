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
	http.HandleFunc("/todos/", handlers.DeleteTodo)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}

}
