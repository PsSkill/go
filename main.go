package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handle(c *gin.Context){
	response:= map[string]interface{}{
		"data":"admin",
	}

	c.JSON(200,response)
}

func main(){
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}           // Add your frontend URL here
	config.AllowMethods = []string{"GET", "POST"} // Allow only GET and POST requests
	config.AllowHeaders = []string{"Authorization", "organisation", "Content-Type"}

	r.Use(cors.New(config))

	r.POST("/sample",handle)


	err := r.Run(":8080")
	if err != nil {
		fmt.Print(err.Error())
	}
}