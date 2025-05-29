package user

import "log"

type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct {
	logger     *log.Logger
	repository Repository
}

func NewService(logger *log.Logger, repository Repository) Service {
	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) Create(firstName, lastName, email, phone string) error {
	s.logger.Println("Create user service")

	var newUser = User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	s.repository.Create(&newUser)
	return nil
}
