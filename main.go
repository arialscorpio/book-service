package main

import (
	"net/http"

	"github.com/arialscorpio/book-service/app"
)

func main() {
	b := app.Book{
		Name:      "TBook",
		Author:    "Jon Doe",
		ISBN:      "12345678",
		Publisher: "XY",
	}

	store := app.NewStore()
	store.Add(b)

	controller := app.NewBookController(&store)

	router := http.NewServeMux()

	router.HandleFunc("/list", controller.List)
	router.HandleFunc("/add", controller.Create)
	router.HandleFunc("/update", controller.Update)
	router.HandleFunc("/delete", controller.Delete)

	_ = http.ListenAndServe(":8088", router)
}
