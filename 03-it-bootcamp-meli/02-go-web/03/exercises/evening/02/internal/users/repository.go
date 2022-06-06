package users

import (
	"fmt"

	"github.com/Luiz-Wendel/go-studies/03-it-bootcamp-meli/02-go-web/03/exercises/evening/02/pkg/store"
)

var users []User = []User{}

type Repository interface {
	GetAll() ([]User, error)
	Store(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error)
	Update(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error)
	PartialUpdate(id, lastName string, age uint) (User, error)
	Delete(id string) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) Store(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error) {
	var us []User

	r.db.Read(&us)
	user := User{id, firstName, lastName, email, creationDate, age, height, active}
	us = append(us, user)

	if err := r.db.Write(us); err != nil {
		return User{}, err
	}

	return user, nil
}

func (repository) Update(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error) {
	user := User{id, firstName, lastName, email, creationDate, age, height, active}
	updated := false

	for index := range users {
		if users[index].Id == id {
			users[index] = user
			updated = true
		}
	}

	if !updated {
		return User{}, fmt.Errorf("user with id: \"%s\" not found", id)
	}

	return user, nil
}

func (repository) PartialUpdate(id, lastName string, age uint) (User, error) {
	var user User
	updated := false

	for index := range users {
		if users[index].Id == id {
			users[index].LastName = lastName
			users[index].Age = age

			user = users[index]
			updated = true
		}
	}

	if !updated {
		return User{}, fmt.Errorf("user with id: \"%s\" not found", id)
	}

	return user, nil
}

func (repository) Delete(id string) error {
	deleted := false
	foundIndex := 0

	for index := range users {
		if users[index].Id == id {
			foundIndex = index
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("user with id: \"%s\" not found", id)
	}

	users = append(users[:foundIndex], users[foundIndex+1:]...)

	return nil
}
