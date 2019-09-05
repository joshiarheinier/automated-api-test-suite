package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {

	// Ping test
	r.GET("/v1/testcase/one", ControlOne)
	r.POST("/v1/testcase/two", ControlTwo)
	r.POST("/v1/testcase/three", ControlThree)
}
