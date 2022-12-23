//A dummy controller that will be secured by JWT Authentication.
//This is just to showcase the ability of the middleware that we will build to restrict access to only
//the requests that have an actual valid JWT in the request header.

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
