package bootstrap

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	CreateDatabaseIfNotExist(app.Env)
	// dsn := "root:meng0987612345@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		app.Env.DBUser, app.Env.DBPass, app.Env.DBHost, app.Env.DBPort, app.Env.DBName)
	var err error
	app.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error cannot open database error: ", err)
	}

	return *app
}

func CreateDatabaseIfNotExist(env *Env) error {
	// 第一步：先连接MySQL服务器（不指定数据库）
	initDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser, env.DBPass, env.DBHost, env.DBPort)

	// 创建临时连接
	tempDB, err := gorm.Open(mysql.Open(initDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("无法连接到MySQL服务器:", err)
		return err
	}

	// 第二步：创建数据库（如果不存在）
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", env.DBName)
	tempDB.Exec(createDBSQL)

	// 第三步：关闭临时连接
	sqlDB, _ := tempDB.DB()
	sqlDB.Close()
	return nil
}
