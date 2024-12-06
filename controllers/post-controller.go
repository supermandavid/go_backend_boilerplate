package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/supermandavid/go_backend_boilerplate/entities"
	"github.com/supermandavid/go_backend_boilerplate/errors"
	"github.com/supermandavid/go_backend_boilerplate/response"
	"github.com/supermandavid/go_backend_boilerplate/services/post"
	"net/http"
)

type controller struct{}

var (
	postService services.PostService
)

type PostController interface {
	GetPosts(ctx *gin.Context)
	AddPost(ctx *gin.Context)
}

func NewPostController(service services.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(ctx *gin.Context) {

	posts, err := postService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.ServiceError{"Cannot find posts"})
		return
	}
	ctx.JSON(http.StatusOK, response.NewHTTPResponse(true, nil, posts))
}

func (*controller) AddPost(ctx *gin.Context) {

	var post entities.Post
	err := json.NewDecoder(ctx.Request.Body).Decode(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.ServiceError{"Error unmarshalling post"})
		return
	}

	validationErr := postService.Validate(&post)

	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, errors.ServiceError{"Error with post data"})
		return
	}

	result, createErr := postService.Create(&post)

	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, errors.ServiceError{"Error saving data"})
		return
	}

	ctx.JSON(http.StatusCreated, response.NewHTTPResponse(true, nil, result))
}
