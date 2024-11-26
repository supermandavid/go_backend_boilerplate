package main

import (
	"awesomeBackend/controllers"
	router "awesomeBackend/http"
	"awesomeBackend/repositories"
	"awesomeBackend/response"
	"awesomeBackend/services"
	"github.com/gin-gonic/gin"
)

var (
	// post service dependencies
	postRepository = repositories.NewFirestoreRepository()
	postService    = services.NewPostService(postRepository)
	postController = controllers.NewPostController(postService)

	carDetailsService    = services.NewCarDetailsService()
	carDetailsController = controllers.NewCarDetailsController(carDetailsService)

	httpRouter = router.NewGinRouter()
)

func main() {
	const port string = ":8082"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, response.NewHTTPResponse("Hello there", nil, "payload"))
	})

	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.Serve(port)

}
