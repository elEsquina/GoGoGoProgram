package data

import (
	"sync"
	"errors"
)

type BookRepository struct {
	mu     	sync.RWMutex
	books  	*map[int]Book
	nextID 	int
}

func NewBookRepository(store *InMemoryStore) *BookRepository {
	highestID := 0
	for id := range store.Books {
		if id > highestID {
			highestID = id
		}
	}
	return &BookRepository{
		books:  &store.Books,
		nextID: highestID + 1,
	}
}


func (repo *BookRepository) Create(book Book) (Book, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	book.ID = repo.nextID
	repo.nextID++
	(*repo.books)[book.ID] = book
	return book, nil
}

func (repo *BookRepository) GetById(id int) (Book, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	book, exists := (*repo.books)[id]
	if !exists {
		return Book{}, errors.New("book not found")
	}
	return book, nil
}

func (repo *BookRepository) Update(id int, updated Book) (Book, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	book, exists := (*repo.books)[id]
	if !exists {
		return Book{}, errors.New("book not found")
	}
	updated.ID = book.ID
	(*repo.books)[id] = updated
	return updated, nil
}

func (repo *BookRepository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, exists := (*repo.books)[id]
	if !exists {
		return errors.New("book not found")
	}
	delete((*repo.books), id)
	return nil
}

func (repo *BookRepository) GetAll() ([]Book, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	books := make([]Book, 0, len(*repo.books))
	for _, book := range *repo.books {
		books = append(books, book)
	}
	return books, nil
}