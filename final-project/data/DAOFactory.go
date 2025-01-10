package data

import "errors"

var m = make(map[string] any)

func GetDAO[T EntityType](name string, storage *InMemoryStore) (IDAO[T], error) {
	if dao, exists := m[name]; exists {
		if typedDAO, ok := dao.(IDAO[T]); ok {
			return typedDAO, nil
		}
		return nil, errors.New("repository exists with a different type")
	}

	var dao IDAO[T]
	switch name {
	case "book":
		if bookRepo, ok := any(NewBookRepository(storage)).(IDAO[T]); ok {
			dao = bookRepo
		}
	case "customer":
		if customerRepo, ok := any(NewCustomerRepository(storage)).(IDAO[T]); ok {
			dao = customerRepo
		}
	case "author":
		if authorRepo, ok := any(NewAuthorRepository(storage)).(IDAO[T]); ok {
			dao = authorRepo
		}
	case "order":
		if orderRepo, ok := any(NewOrderRepository(storage)).(IDAO[T]); ok {
			dao = orderRepo
		}
	default:
		return nil, errors.New("invalid repository")
	}

	if dao == nil {
		return nil, errors.New("failed to create repository")
	}

	m[name] = dao
	return dao, nil
}
