package data

import (
	"errors"
	"sync"
)

type CustomerRepository struct {
	mu        sync.RWMutex
	customers *map[int]Customer
	nextID    int
}

func NewCustomerRepository(store *InMemoryStore) *CustomerRepository {
	highestID := 0
	for id := range store.Customers {
		if id > highestID {
			highestID = id
		}
	}
	return &CustomerRepository{
		customers: &store.Customers,
		nextID:    highestID + 1,
	}
}

func (repo *CustomerRepository) Create(customer Customer) (Customer, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	customer.ID = repo.nextID
	repo.nextID++
	(*repo.customers)[customer.ID] = customer
	return customer, nil
}

func (repo *CustomerRepository) GetById(id int) (Customer, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	customer, exists := (*repo.customers)[id]
	if !exists {
		return Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func (repo *CustomerRepository) Update(id int, updated Customer) (Customer, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	customer, exists := (*repo.customers)[id]
	if !exists {
		return Customer{}, errors.New("customer not found")
	}
	updated.ID = customer.ID
	(*repo.customers)[id] = updated
	return updated, nil
}

func (repo *CustomerRepository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, exists := (*repo.customers)[id]
	if !exists {
		return errors.New("customer not found")
	}
	delete((*repo.customers), id)
	return nil
}

func (repo *CustomerRepository) GetAll() ([]Customer, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	customers := make([]Customer, 0, len(*repo.customers))
	for _, customer := range *repo.customers {
		customers = append(customers, customer)
	}
	return customers, nil
}
