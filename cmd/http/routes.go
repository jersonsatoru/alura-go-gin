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
	router.GET("/students/:id", studentHandler.getStudentByIDHandler)
	router.POST("/students", studentHandler.createStudentHandler)
	router.DELETE("/students/:id", studentHandler.deleteStudentHandler)
	router.PUT("/students/:id", studentHandler.updateStudentHandler)
	router.GET("/students/cpf/:cpf", studentHandler.findStudentByCPFHandler)
	return router
}
