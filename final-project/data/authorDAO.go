package data

import (
	"errors"
	"sync"
)

type AuthorRepository struct {
	mu      sync.RWMutex
	authors *map[int]Author
	nextID  int
}

func NewAuthorRepository(store *InMemoryStore) *AuthorRepository {
	highestID := 0
	for id := range store.Authors {
		if id > highestID {
			highestID = id
		}
	}
	return &AuthorRepository{
		authors: &store.Authors,
		nextID:  highestID + 1,
	}
}

func (repo *AuthorRepository) Create(author Author) (Author, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	author.ID = repo.nextID
	repo.nextID++
	(*repo.authors)[author.ID] = author
	return author, nil
}

func (repo *AuthorRepository) GetById(id int) (Author, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	author, exists := (*repo.authors)[id]
	if !exists {
		return Author{}, errors.New("author not found")
	}
	return author, nil
}

func (repo *AuthorRepository) Update(id int, updated Author) (Author, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	author, exists := (*repo.authors)[id]
	if !exists {
		return Author{}, errors.New("author not found")
	}
	updated.ID = author.ID
	(*repo.authors)[id] = updated
	return updated, nil
}

func (repo *AuthorRepository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, exists := (*repo.authors)[id]
	if !exists {
		return errors.New("author not found")
	}
	delete((*repo.authors), id)
	return nil
}

func (repo *AuthorRepository) GetAll() ([]Author, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	authors := make([]Author, 0, len(*repo.authors))
	for _, author := range *repo.authors {
		authors = append(authors, author)
	}
	return authors, nil
}