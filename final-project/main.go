package main

import (
	"net/http"
	"finalproject/data"
	"finalproject/api"
)

// Saving, Greaceful Shutdown, and Readme, Saving from update and create and delete, postgres

func main() {
	store := data.NewInMemoryStore()
	//http.Handle("/login", http.HandlerFunc(api.Login))

	go data.StartReportGenerator(store)

	http.Handle("/books", 
		api.RequestLogger(
			api.ContextGeneration(store, http.HandlerFunc(api.BooksRouter)),
		),
	)

	http.Handle("/books/{id}", 
		api.RequestLogger(
			api.ContextGeneration(store, http.HandlerFunc(api.BooksPathParamRouter)),
		),
	)

	http.Handle("/authors", 
		api.RequestLogger(
			api.ContextGeneration(store, http.HandlerFunc(api.AuthorsRouter)),
		),
	)

	http.Handle("/authors/{id}", 
		api.RequestLogger(
			api.ContextGeneration(store, http.HandlerFunc(api.AuthorsPathParamRouter)),
		),
	)

	http.ListenAndServe(":8080", nil)
}

