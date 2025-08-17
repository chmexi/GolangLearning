package usecase

import (
	"context"
	"time"

	"github.com/chmexi/GolangLearning/task4/domain"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
	contextTimeOut    time.Duration
}

func NewCommentUsecase(commentRepository domain.CommentRepository, timeout time.Duration) domain.CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		contextTimeOut:    timeout,
	}
}

func (cu *commentUsecase) Create(c context.Context, comment domain.Comment) error {
	return cu.commentRepository.Create(c, comment)
}

func (cu *commentUsecase) GetCommentsByPostID(c context.Context, postID int) ([]domain.Comment, error) {
	return cu.commentRepository.GetCommentsByPostID(c, postID)
}
