package usecase

import (
	"context"
	"time"

	"github.com/chmexi/GolangLearning/task4/domain"
	tokenutil "github.com/chmexi/GolangLearning/task4/internal/token_util"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	return su.userRepository.Create(c, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	return su.userRepository.GetByEmail(c, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}
func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
