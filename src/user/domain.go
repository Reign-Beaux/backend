package user

import "time"

type User struct {
	ID        string     `json:"id" gorm:"type: char(36);not null;primary_key: unique_index"`
	FirstName string     `json:"first_name" gorm:"type: char(50);not null"`
	LastName  string     `json:"last_name" gorm:"type: char(50);not null"`
	Email     string     `json:"email" gorm:"type: char(50);not null"`
	Phone     string     `json:"phone" gorm:"type: char(10);not null"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
