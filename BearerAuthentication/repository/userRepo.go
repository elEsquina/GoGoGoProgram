package repository

import (
	"encoding/json"
	"io"
	"os"
	"bearerAuth/model"
	"errors"
	"fmt"
) 

type UserRepository struct {
	Users []model.User
}

func NewUserRepository() *UserRepository {
	m := &UserRepository{
		Users: []model.User{},
	}
	
	fmt.Println(m.Load("users.json"))
	return m
}


func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	for _, u := range r.Users {
		if u.Email == email {
			return u, nil
		}
	}

	return model.User{}, errors.New("User not found")
}

func (r *UserRepository) Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &r.Users)
	return err
}


