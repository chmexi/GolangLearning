package controller

import (
	"net/http"
	"strconv"

	"github.com/chmexi/GolangLearning/task4/bootstrap"
	"github.com/chmexi/GolangLearning/task4/domain"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostUsecase domain.PostUsecase
	Env         *bootstrap.Env
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	posts, err := pc.PostUsecase.GetAllPosts(c)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "invalid post id"})
		return
	}
	post, _ := pc.PostUsecase.GetPostByID(c, id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "指定文章不存在"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *PostController) CreatePost(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("x-user-id"))
	var request domain.CreatePostRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.PostUsecase.Create(c, domain.Post{
		UserID:  userID,
		Title:   request.Title,
		Content: request.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "创建文章成功"})
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("x-user-id"))
	var request domain.UpdatePostRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	post, _ := pc.PostUsecase.GetPostByID(c, request.PostID)
	if userID != post.UserID {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "you cannot delete the post that is not created by you"})
		return
	}

	post = domain.Post{
		UserID:  userID,
		Title:   request.Title,
		Content: request.Content,
	}
	post.ID = uint(request.PostID)
	err := pc.PostUsecase.UpdatePost(c, post)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "update successfully"})
}

func (pc *PostController) DeletePost(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("x-user-id"))
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	post, _ := pc.PostUsecase.GetPostByID(c, postID)
	if userID != post.UserID {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "you cannot delete the post that is not created by you"})
		return
	}

	err = pc.PostUsecase.DeletePost(c, postID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "you have deleted the post successfully"})
}
