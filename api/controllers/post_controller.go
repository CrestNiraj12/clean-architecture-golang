package controllers

import (
	"clean-architecture/lib"
	"clean-architecture/models"
	"clean-architecture/repository"
	"clean-architecture/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service    services.PostService
	logger     lib.Logger
	repository repository.PostRepository
}

func NewPostController(
	postService services.PostService,
	logger lib.Logger,
	postRepository repository.PostRepository,
) PostController {
	return PostController{
		service:    postService,
		logger:     logger,
		repository: postRepository,
	}
}

func (c PostController) GetOnePost(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	post, err := c.service.GetOnePost(uint(id))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func (c PostController) GetPosts(ctx *gin.Context) {
	posts, err := c.service.GetAllPosts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func (c PostController) SavePost(ctx *gin.Context) {
	var post models.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := c.service.SavePost(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Post created",
	})
}

func (c PostController) UpdatePost(ctx *gin.Context) {
	post := models.Post{}
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err = c.service.UpdatePost(uint(id), post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Post updated",
	})
}

func (c PostController) DeletePost(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err = c.service.DeletePost(uint(id)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Post deleted",
	})
}
