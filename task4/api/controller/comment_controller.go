package controller

import (
	"net/http"
	"strconv"

	"github.com/chmexi/GolangLearning/task4/bootstrap"
	"github.com/chmexi/GolangLearning/task4/domain"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentUsecase domain.CommentUsecase
	Env            *bootstrap.Env
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("x-user-id"))

	var comment domain.Comment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	comment.UserID = userID
	err := cc.CommentUsecase.Create(c, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "you have created comment successfully"})
}

func (cc *CommentController) GetCommentsByPostID(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	comments, err := cc.CommentUsecase.GetCommentsByPostID(c, postID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, comments)
}
