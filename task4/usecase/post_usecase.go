package usecase

import (
	"context"
	"time"

	"github.com/chmexi/GolangLearning/task4/domain"
)

type postUseCase struct {
	postRepository domain.PostRepository
	contextTimeOut time.Duration
}

func NewPostUseCase(postRepository domain.PostRepository, timeout time.Duration) domain.PostUsecase {
	return &postUseCase{
		postRepository: postRepository,
		contextTimeOut: timeout,
	}
}

func (pu *postUseCase) Create(c context.Context, post domain.Post) error {
	return pu.postRepository.Create(c, post)
}

func (pu *postUseCase) GetAllPosts(c context.Context) ([]domain.Post, error) {
	return pu.postRepository.GetAllPosts(c)
}

func (pu *postUseCase) GetPostByID(c context.Context, postID int) (domain.Post, error) {
	return pu.postRepository.GetPostByID(c, postID)
}

func (pu *postUseCase) UpdatePost(c context.Context, post domain.Post) error {
	return pu.postRepository.UpdatePost(c, post)
}

func (pu *postUseCase) DeletePost(c context.Context, postID int) error {
	return pu.postRepository.DeletePost(c, postID)
}
