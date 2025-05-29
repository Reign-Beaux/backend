package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
}

type repo struct {
	logger *log.Logger
	db     *gorm.DB
}

func NewRepository(logger *log.Logger, db *gorm.DB) Repository {
	return &repo{
		logger: logger,
		db:     db,
	}
}

func (r *repo) Create(user *User) error {
	user.ID = uuid.New().String()
	r.logger.Println("Creating user")
	var result = r.db.Create(user)

	if result.Error != nil {
		r.logger.Printf("Error creating user: %v", result.Error)
		return result.Error
	}

	return nil
}
