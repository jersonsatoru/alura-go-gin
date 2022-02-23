package main

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.GET("/students/:name", listStudentsHandler)
	return router
}
