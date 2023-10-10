package app

import "golang.org/x/exp/slices"

type (
	Book struct {
		Name      string `json:"name"`
		Author    string `json:"author"`
		ISBN      string `json:"isbn"`
		Publisher string `json:"publisher"`
	}

	// Store defines a collection or storage of books.
	Store []Book
)

func NewStore() Store {
	return make(Store, 0)
}

func (s *Store) GetAll() []Book {
	return []Book(*s)
}

// Add adds a new book to the store.
func (s *Store) Add(b Book) {
	for _, book := range *s {
		if book.ISBN == b.ISBN {
			return
		}
	}

	*s = append(*s, b)
}

// Updates an existing book in the store.
func (s *Store) Update(b Book) {
	for i, book := range *s {
		if book.ISBN == b.ISBN {
			(*s)[i] = b
			return
		}
	}
}

// Delete deletes the book with given ISBN from the store.
func (s *Store) Delete(ISBN string) {
	for i, book := range *s {
		if book.ISBN == ISBN {
			*s = slices.Delete(*s, i, i+1)
			return
		}
	}
}
