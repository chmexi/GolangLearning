package route

import (
	"time"

	"github.com/chmexi/GolangLearning/task4/api/controller"
	"github.com/chmexi/GolangLearning/task4/bootstrap"
	"github.com/chmexi/GolangLearning/task4/repository"
	"github.com/chmexi/GolangLearning/task4/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUseCase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
