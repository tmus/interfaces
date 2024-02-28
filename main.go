package main

import (
	"github.com/gin-gonic/gin"
)

type Habit struct {
	Title string
}

type HabitRepository interface {
	GetAll() []Habit
}

func main() {
	repository := NewInMemoryHabitRepository()

	habitHandler := NewHabitHandler(repository)

	router := gin.Default()

	router.GET("/habits", habitHandler.GetHabits)

	router.Run(":8080")
}
