package main

import (
	"net/http"

	"github.com/arialscorpio/book-service/app"
	"github.com/arialscorpio/ms-lib/logger"
)

type Middleware func(http.Handler) http.Handler

func main() {
	logger := logger.New()
	b := app.Book{
		Name:      "TBook",
		Author:    "Jon Doe",
		ISBN:      "12345678",
		Publisher: "XY",
	}

	store := app.NewStore()
	store.Add(b)

	controller := app.NewBookController(&store, logger)

	router := http.NewServeMux()

	router.HandleFunc("/list", controller.List)
	router.HandleFunc("/add", controller.Create)
	router.HandleFunc("/update", controller.Update)
	router.HandleFunc("/delete", controller.Delete)

	addr := ":8088"
	logger.Info("starting server on " + addr)
	_ = http.ListenAndServe(addr, router)
}
