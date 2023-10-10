package app

import (
	"encoding/json"
	"io"
	"net/http"
)

type BookController struct {
	store *Store
}

func NewBookController(s *Store) *BookController {
	return &BookController{s}
}

func (c *BookController) List(w http.ResponseWriter, r *http.Request) {
	books := c.store.GetAll()
	res, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *BookController) Create(w http.ResponseWriter, r *http.Request) {
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

	var b Book
	if err := json.Unmarshal(rp, &b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.store.Add(b)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *BookController) Update(w http.ResponseWriter, r *http.Request) {
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

	var b Book
	if err := json.Unmarshal(rp, &b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.store.Update(b)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isbn := r.URL.Query().Get("isbn")
	if isbn == "" {
		return
	}

	c.store.Delete(isbn)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
