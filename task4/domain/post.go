package domain

import (
	"context"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID        int
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Comments      []Comment `json:"comments" gorm:"foreignKey:PostID;OnDelete:CASCADE"`
	CommentCount  int       `json:"commentCount"`
	CommentStatus string    `json:"commentStatus" gorm:"default:'无评论'" `
}

type CreatePostRequest struct {
	Title       string `form:"title" binding:"required"`
	Content     string `form:"content" binding:"required"`
	AccessToken string `form:"accessToken" binding:"required"`
}

type UpdatePostRequest struct {
	PostID      int    `form:"postID" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Content     string `form:"content" binding:"required"`
	AccessToken string `form:"accessToken" binding:"required"`
}

type PostRepository interface {
	Create(c context.Context, post Post) error
	GetPostByID(c context.Context, postID int) (Post, error)
	GetAllPosts(c context.Context) ([]Post, error)
	UpdatePost(c context.Context, post Post) error
	DeletePost(c context.Context, postID int) error
}

type PostUsecase interface {
	Create(c context.Context, post Post) error
	GetPostByID(c context.Context, postID int) (Post, error)
	GetAllPosts(c context.Context) ([]Post, error)
	UpdatePost(c context.Context, post Post) error
	DeletePost(c context.Context, postID int) error
}
