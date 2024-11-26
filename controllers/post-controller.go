package controllers

import (
	"awesomeBackend/entities"
	"awesomeBackend/errors"
	"awesomeBackend/response"
	"awesomeBackend/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

	fmt.Println("I'm here")
	posts, err := postService.FindAll()
	if err != nil {
		fmt.Println("hit error")
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, errors.ServiceError{"Cannot find posts"})
		return
	}
	fmt.Println("finsished")
	fmt.Println(len(posts))
	ctx.JSON(http.StatusOK, posts)
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

	ctx.JSON(http.StatusOK, response.NewHTTPResponse(nil, nil, result))
}
