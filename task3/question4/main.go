package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

type User struct {
	gorm.Model
	Name      string
	Account   string
	Password  string
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount uint
}

type Post struct {
	gorm.Model
	UserID        uint
	Title         string
	Content       string
	Comments      []Comment `gorm:"foreignKey:PostID;OnDelete:CASCADE"`
	CommentCount  uint
	CommentStatus string `gorm:"default:'无评论'"`
}

type Comment struct {
	gorm.Model
	PostID  uint
	Content string
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	var user User
	tx.Find(&user, p.UserID)
	user.PostCount++
	tx.Model(&user).Update("post_count", user.PostCount)
	return nil
}

func (p *Post) AfterDelete(tx *gorm.DB) error {
	var user User
	tx.Find(&user, p.UserID)
	user.PostCount--
	tx.Model(&user).Update("post_count", user.PostCount)
	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	postOfComment := Post{}
	result := tx.Where("id = ?", c.PostID).Find(&postOfComment)
	if result.Error != nil {
		fmt.Println("查找评论所属文章失败 error:", result.Error)
		return result.Error
	}
	postOfComment.CommentCount++
	postOfComment.CommentStatus = "有评论"
	result = tx.Save(&postOfComment)
	if result.Error != nil {
		fmt.Println("更新文章评论数量 error:", result.Error)
		return result.Error
	}
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	postOfComment := Post{}
	if result := tx.Where("id = ?", c.PostID).Find(&postOfComment); result.Error != nil {
		fmt.Println("查找评论所属文章失败 error:", result.Error)
		return result.Error
	}
	postOfComment.CommentCount--
	if postOfComment.CommentCount == 0 {
		postOfComment.CommentStatus = "无评论"
	}

	if result := tx.Save(&postOfComment); result.Error != nil {
		fmt.Println("更新文章评论状态及数量失败 error:", result.Error)
		return result.Error
	}
	return nil
}

func InsertTestData(db *gorm.DB) {
	// 创建用户
	user := User{
		Name:     "John Doe",
		Account:  "johndoe",
		Password: "password123",
		Posts: []Post{
			{
				Title:   "First Post",
				Content: "This is the content of the first post.",
				Comments: []Comment{
					{
						Content: "Great post!",
					},
					{
						Content: "Very informative.",
					},
				},
			},
			{
				Title:   "Second Post",
				Content: "This is the content of the second post.",
				Comments: []Comment{
					{
						Content: "Nice one!",
					},
				},
			},
		},
	}

	// 插入用户及其关联的帖子和评论
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Printf("插入数据失败: %v\n", result.Error)
		return
	}
}

func GetAllPostsAndCommentByUserName(db *gorm.DB, userName string) {
	var user User
	result := db.Preload("Posts").Preload("Posts.Comments").Where("name = ?", userName).Find(&user)
	if result.Error != nil {
		fmt.Println("查询用户文章及评论错误 error:", result.Error)
	}
	fmt.Println(user)
}

func GetPostWithMaxCommentCount(db *gorm.DB) Post {
	post := Post{}
	db.Order("comment_count DESC").Find(&post)
	return post
}

func main() {
	dsn := "root:meng0987612345@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error cannot open database error: ", err)
	}
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("创建表失败 error: ", err)
	}
	// InsertTestData(db)
	// GetAllPostsAndCommentByUserName(db, "John Doe")
	// post := GetPostWithMaxCommentCount(db)
	// fmt.Println(post)

	// var comment Comment
	// db.Find(&comment, 3)
	// db.Delete(&comment)

	var post Post
	db.Find(&post, 2)
	db.Delete(&post)
}
