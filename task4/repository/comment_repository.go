package repository

import (
	"context"

	"github.com/chmexi/GolangLearning/task4/domain"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	db.AutoMigrate(&domain.Comment{})
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) Create(c context.Context, comment domain.Comment) error {
	result := cr.db.WithContext(c).Create(&comment)
	return result.Error
}

func (cr *commentRepository) GetCommentsByPostID(c context.Context, postID int) ([]domain.Comment, error) {
	var comments []domain.Comment
	result := cr.db.WithContext(c).Where("post_id = ?", postID).Find(&comments)
	return comments, result.Error
}
