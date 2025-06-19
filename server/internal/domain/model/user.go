package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email string) *User {
	now := time.Now()
	id, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate UUID: " + err.Error())
	}
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
