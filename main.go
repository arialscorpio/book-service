package main

import (
	"encoding/json"
	"io"
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

	router := http.NewServeMux()

	router.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		books := store.GetAll()
		res, err := json.Marshal(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	router.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer r.Body.Close()
		rp, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var b app.Book
		if err := json.Unmarshal(rp, &b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		store.Add(b)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return
	})

	router.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer r.Body.Close()
		rp, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var b app.Book
		if err := json.Unmarshal(rp, &b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		store.Update(b)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	router.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		isbn := r.URL.Query().Get("isbn")
		if isbn == "" {
			return
		}

		store.Delete(isbn)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	_ = http.ListenAndServe(":8088", router)
}
