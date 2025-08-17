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

func NewCommentRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	cr := repository.NewCommentRepository(db)
	cc := controller.CommentController{
		CommentUsecase: usecase.NewCommentUsecase(cr, timeout),
		Env:            env,
	}
	group.POST("/comment", cc.CreateComment)
	group.GET("/comment/:post_id", cc.GetCommentsByPostID)
}
