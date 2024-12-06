package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ginRouter struct {
}

var (
	ginDispatcher = gin.New()
)

func NewGinRouter() Router {

	// Open or create the log file
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Set Gin's default writer to both file and stdout
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	absPath, err := filepath.Abs("gin.log")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Absolute file path: %s\n", absPath)

	ginDispatcher.Use(gin.Recovery(), gin.Logger())
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func(*gin.Context)) {
	ginDispatcher.GET(uri, f)
}

func (*ginRouter) POST(uri string, f func(*gin.Context)) {
	ginDispatcher.POST(uri, f)
}

func (*ginRouter) Serve(port string) {
	fmt.Printf("Listening on port %s\n", port)
	ginDispatcher.Run(":" + port)
}

func (*ginRouter) ServeRequest(w http.ResponseWriter, req *http.Request) {
	ginDispatcher.ServeHTTP(w, req)
}
