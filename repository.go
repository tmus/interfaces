package main

type InMemoryHabitRepository struct {
	habits []Habit
}

func NewInMemoryHabitRepository() *InMemoryHabitRepository {
	return &InMemoryHabitRepository{
		habits: []Habit{
			Habit{Title: "Read book"},
			Habit{Title: "Wash dishes"},
		},
	}
}
