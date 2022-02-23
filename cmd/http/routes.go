package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jersonsatoru/alura-go-gin/internal/repositories"
)

func Routes() *gin.Engine {
	router := gin.Default()
	repository := &repositories.StudentRepository{}
	studentHandler := NewStudentHandler(repository)
	router.GET("/students", studentHandler.listStudentsHandler)
	router.POST("/students", studentHandler.createStudentHandler)
	return router
}
