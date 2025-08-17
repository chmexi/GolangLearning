package repository

import (
	"context"

	"github.com/chmexi/GolangLearning/task4/domain"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	db.AutoMigrate(&domain.Post{})
	return &postRepository{
		db: db,
	}
}

func (pr *postRepository) Create(c context.Context, post domain.Post) error {
	result := pr.db.WithContext(c).Create(&post)
	return result.Error
}

func (pr *postRepository) GetAllPosts(c context.Context) ([]domain.Post, error) {
	var posts []domain.Post
	result := pr.db.WithContext(c).Find(&posts)
	return posts, result.Error
}

func (pr *postRepository) GetPostByID(c context.Context, postID int) (domain.Post, error) {
	var post domain.Post
	result := pr.db.WithContext(c).Where("id = ?", postID).Find(&post)
	return post, result.Error
}
func (pr *postRepository) UpdatePost(c context.Context, post domain.Post) error {
	result := pr.db.WithContext(c).Updates(&post)
	return result.Error
}

func (pr *postRepository) DeletePost(c context.Context, postID int) error {
	result := pr.db.WithContext(c).Where("id = ?", postID).Delete(&domain.Post{})
	return result.Error
}
