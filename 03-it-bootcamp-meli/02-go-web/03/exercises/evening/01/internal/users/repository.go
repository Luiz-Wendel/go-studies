package users

import (
	"fmt"
)

var users []User = []User{}

type Repository interface {
	GetAll() ([]User, error)
	Store(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error)
	Update(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error)
	PartialUpdate(id, lastName string, age uint) (User, error)
	Delete(id string) error
}

type repository struct{}

func (repository) GetAll() ([]User, error) {
	return users, nil
}

func (repository) Store(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error) {
	user := User{id, firstName, lastName, email, creationDate, age, height, active}

	users = append(users, user)

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

func NewRepository() Repository {
	return &repository{}
}
