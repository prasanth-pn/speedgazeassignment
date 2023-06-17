package main

import (
	"assignment2/handler"

	"github.com/gin-gonic/gin"
)

func main() {


	server := gin.Default()

	
	server.GET("/", handler.GetBreedByCountry)
	server.POST("/", handler.PayloadCheck)

	err := server.Run(":3030")

	if err != nil {
		panic(err)
	}
}
