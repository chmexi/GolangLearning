package usecase

import (
	"context"
	"time"

	"github.com/chmexi/GolangLearning/task4/domain"
	tokenutil "github.com/chmexi/GolangLearning/task4/internal/token_util"
)

type loginUseCase struct {
	userRepository domain.UserRepository
	contextTimeOut time.Duration
}

func NewLoginUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUseCase{
		userRepository: userRepository,
		contextTimeOut: timeout,
	}
}

// GetUserByEmail(c context.Context, email string) (User, error)
// CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
// CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
func (lu *loginUseCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	user, err := lu.userRepository.GetByEmail(c, email)
	return user, err
}

func (lu *loginUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
