package domain

import (
	"context"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  int
	PostID  int    `json:"postID" form:"postID" binding:"required"`
	Content string `json:"comment" form:"comment" binding:"required"`
}

type CommentUsecase interface {
	Create(c context.Context, comment Comment) error
	GetCommentsByPostID(c context.Context, postID int) ([]Comment, error)
}

type CommentRepository interface {
	Create(c context.Context, comment Comment) error
	GetCommentsByPostID(c context.Context, postID int) ([]Comment, error)
}
