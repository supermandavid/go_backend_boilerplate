package main

import (
	"github.com/supermandavid/go_backend_boilerplate/controllers"
	router "github.com/supermandavid/go_backend_boilerplate/http"
	postRepo "github.com/supermandavid/go_backend_boilerplate/repositories/post"
	carSrv "github.com/supermandavid/go_backend_boilerplate/services/car"
	postSrv "github.com/supermandavid/go_backend_boilerplate/services/post"
	"os"
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

	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.Serve(os.Getenv("PORT"))

}
