package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jersonsatoru/alura-go-gin/internal/entities"
	"github.com/jersonsatoru/alura-go-gin/internal/repositories"
	"gorm.io/gorm"
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
	students, err := sh.studentRepository.Find()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
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

func (sh *StudentHandler) getStudentByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	student, err := sh.studentRepository.FindByID(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("student with id %d was not found", id),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, student)
}

func (s *StudentHandler) deleteStudentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = s.studentRepository.Delete(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "cannot delete a non existent student",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
	c.JSON(http.StatusNoContent, nil)
}

func (s StudentHandler) updateStudentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var input struct {
		Name string `json:"name"`
		CPF  string `json:"cpf"`
		RG   string `json:"rg"`
	}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	student, err := s.studentRepository.FindByID(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("student with id %d not found", id),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	student.Name = input.Name
	student.CPF = input.CPF
	student.RG = input.RG
	updatedStudent, err := s.studentRepository.Update(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updatedStudent)
}
