package main

import (
	"awesomeBackend/controllers"
	router "awesomeBackend/http"
	postRepo "awesomeBackend/repositories/post"
	"awesomeBackend/response"
	carSrv "awesomeBackend/services/car"
	postSrv "awesomeBackend/services/post"
	"github.com/gin-gonic/gin"
)

var (
	// post service dependencies
	postRepository = postRepo.NewSQLiteRepository()
	postService    = postSrv.NewPostService(postRepository)
	postController = controllers.NewPostController(postService)

	carDetailsService    = carSrv.NewCarDetailsService()
	carDetailsController = controllers.NewCarDetailsController(carDetailsService)

	httpRouter = router.NewGinRouter()
)

func main() {
	const port string = ":8082"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, response.NewHTTPResponse(true, "payload", "done"))
	})

	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.Serve(port)

}
