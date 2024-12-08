package entity

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/bianavic/fullcycle_apis/pkg/entity"
)

// entity - regras de negocio
type User struct {
	ID       entity.ID `json:"ID"`
	Name     string    `json:"Name"`
	Email    string    `json:"Email"`
	Password string    `json:"Password"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
