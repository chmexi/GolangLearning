package route

import (
	"time"

	"github.com/chmexi/GolangLearning/task4/api/middleware"
	"github.com/chmexi/GolangLearning/task4/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUp(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// register all public APIs
	NewLoginRouter(env, timeout, db, publicRouter)
	NewSignupRouter(env, timeout, db, publicRouter)

	// register all protected APIs
	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	NewPostRouter(env, timeout, db, protectedRouter)
	NewCommentRouter(env, timeout, db, protectedRouter)
}
