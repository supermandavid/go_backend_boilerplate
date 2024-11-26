package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	GET(uri string, f func(ctx *gin.Context))
	POST(uri string, f func(ctx *gin.Context))
	Serve(port string)
}