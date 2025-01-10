package data

import (
	"errors"
	"sync"
)

type OrderRepository struct {
	mu     sync.RWMutex
	orders *map[int]Order
	nextID int
}

func NewOrderRepository(store *InMemoryStore) *OrderRepository {
	highestID := 0
	for id := range store.Orders {
		if id > highestID {
			highestID = id
		}
	}
	return &OrderRepository{
		orders: &store.Orders,
		nextID: highestID + 1,
	}
}

func (repo *OrderRepository) Create(order Order) (Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	order.ID = repo.nextID
	repo.nextID++
	(*repo.orders)[order.ID] = order
	return order, nil
}

func (repo *OrderRepository) GetById(id int) (Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	order, exists := (*repo.orders)[id]
	if !exists {
		return Order{}, errors.New("order not found")
	}
	return order, nil
}

func (repo *OrderRepository) Update(id int, updated Order) (Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	order, exists := (*repo.orders)[id]
	if !exists {
		return Order{}, errors.New("order not found")
	}
	updated.ID = order.ID
	(*repo.orders)[id] = updated
	return updated, nil
}

func (repo *OrderRepository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, exists := (*repo.orders)[id]
	if !exists {
		return errors.New("order not found")
	}
	delete((*repo.orders), id)
	return nil
}

func (repo *OrderRepository) GetByCustomerID(customerID int) ([]Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var orders []Order
	for _, order := range *repo.orders {
		if order.Customer.ID == customerID {
			orders = append(orders, order)
		}
	}
	return orders, nil
}


func (repo *OrderRepository) GetAll() ([]Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	orders := make([]Order, 0, len(*repo.orders))
	for _, order := range *repo.orders {
		orders = append(orders, order)
	}
	return orders, nil
}
