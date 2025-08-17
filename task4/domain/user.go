package domain

import (
	"context"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount uint
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByID(c context.Context, id int) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
}
