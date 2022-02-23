package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entities "github.com/jersonsatoru/alura-go-gin/internal/entites"
)

func listStudentsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": entities.Students,
	})
}
