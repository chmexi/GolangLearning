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

func NewPostRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	pr := repository.NewPostRepository(db)
	pc := &controller.PostController{
		PostUsecase: usecase.NewPostUseCase(pr, timeout),
		Env:         env,
	}
	group.GET("/post", pc.GetAllPosts)
	group.GET("/post/:id", pc.GetPostByID)
	group.POST("/post", pc.CreatePost)
	group.PUT("/post", pc.UpdatePost)
	group.DELETE("/post/:id", pc.DeletePost)
}
