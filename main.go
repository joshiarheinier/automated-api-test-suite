package main

import (
	// "database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joshia/automated-api-test-suite/controllers"
)


func main()  {
	r := gin.Default()

	controllers.SetUpRouter(r)

	// Listen and Server in 0.0.0.0:8182
	r.Run(":8182")
}