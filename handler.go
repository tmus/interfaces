package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HabitHandler struct {
	repository HabitRepository
}

func NewHabitHandler(repository HabitRepository) *HabitHandler {
	return &HabitHandler{
		repository: repository,
	}
}

func (h *HabitHandler) GetHabits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.repository.GetAll())
}
