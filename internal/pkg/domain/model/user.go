package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key;column:id"`
	FullName string    `gorm:"column:full_name"`
	Password string    `gorm:"column:password"`
	Email    string    `gorm:"column:email"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"updated_at;autoCreateTime;autoUpdateTime"`
}

func CreateNewUser(fullName, password string, email string) User {
	return User{
		ID:        uuid.New(),
		FullName:  fullName,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
