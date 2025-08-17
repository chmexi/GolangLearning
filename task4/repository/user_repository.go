package repository

import (
	"context"

	"github.com/chmexi/GolangLearning/task4/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	db.AutoMigrate(&domain.User{})
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	result := ur.db.WithContext(c).Create(user)
	return result.Error
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	result := ur.db.WithContext(c).Find(&users)
	return users, result.Error
}

func (ur *userRepository) GetByID(c context.Context, id int) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(c).Where("id = ?", id).Find(&user)
	return user, result.Error
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	result := ur.db.WithContext(c).Where("email = ?", email).Find(&user)
	return user, result.Error
}
