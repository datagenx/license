package main

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/datagenx/license-generator/controllers"
	"github.com/datagenx/license-generator/internal/initializers"
)

func main() {

	// initilization
	runtime.GOMAXPROCS(runtime.NumCPU())
	st := initializers.Setup()

	if strings.ToUpper(st.Env) == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	//

	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/", welcome)
	router.POST("/generate", controllers.PostGenerate)
	router.POST("/validate", controllers.PostValidate)

	router.Run()
}

// welcome msg
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to License-GO",
	})
}
