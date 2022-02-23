package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jersonsatoru/alura-go-gin/internal/entities"
	"github.com/jersonsatoru/alura-go-gin/internal/repositories"
)

type StudentHandler struct {
	studentRepository repositories.IStudentRepository
}

func NewStudentHandler(repository repositories.IStudentRepository) *StudentHandler {
	return &StudentHandler{
		studentRepository: repository,
	}
}

func (sh *StudentHandler) listStudentsHandler(c *gin.Context) {
	students := sh.studentRepository.Find()
	c.JSON(http.StatusOK, gin.H{
		"message": students,
	})
}

func (sh *StudentHandler) createStudentHandler(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
		RG   string `json:"rg"`
		CPF  string `json:"cpf"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recentlyCreatedStudent, err := sh.studentRepository.Create(&entities.Student{
		Name: input.Name,
		CPF:  input.CPF,
		RG:   input.RG,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Writer.Header().Set("Location", fmt.Sprintf("/students/%d", recentlyCreatedStudent.ID))
	c.JSON(http.StatusCreated, recentlyCreatedStudent)
}
