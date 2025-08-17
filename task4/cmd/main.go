package main

import (
	"time"

	route "github.com/chmexi/GolangLearning/task4/api/route"
	"github.com/chmexi/GolangLearning/task4/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.SetUp(env, timeout, app.DB, gin)

	gin.Run(env.ServerAddress)
}
