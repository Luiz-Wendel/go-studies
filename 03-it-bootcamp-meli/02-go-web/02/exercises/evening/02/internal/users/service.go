package users

import "github.com/google/uuid"

type Service interface {
	GetAll() ([]User, error)
	Store(firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error)
	Update(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s service) Store(firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error) {
	id := uuid.NewString()

	user, err := s.repository.Store(id, firstName, lastName, email, creationDate, age, height, active)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) Update(id, firstName, lastName, email, creationDate string, age, height uint, active bool) (User, error) {
	user, err := s.repository.Update(id, firstName, lastName, email, creationDate, age, height, active)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
