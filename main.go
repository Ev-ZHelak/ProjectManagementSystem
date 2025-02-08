package main

import (
	"courses/project"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	// Создаем проект
	pr, err := project.New(uuid.New(), "Проект 1")
	if err != nil {
		log.Fatalf("create project error: %v\n", err)
	}

	// Создаем задачу 1 для проекта 1
	tk1, err := project.NewTask(uuid.New(), "Задача 1", "Описание важной задачи №1")
	if err != nil {
		log.Fatalf("create task error: %v\n", err)
	}
	// Добавляем задачу в проект 1
	if err := pr.AddTask(*tk1); err != nil {
		log.Fatalf("add task to project error: %v\n", err)
	}

	// Создаем задачу 2 для проекта 1
	tk2, err := project.NewTask(uuid.New(), "Задача 2", "Описание важной задачи №2")
	if err != nil {
		log.Fatalf("create task error: %v\n", err)
	}
	// Добавляем задачу в проект 1
	if err := pr.AddTask(*tk2); err != nil {
		log.Fatalf("add task to project error: %v\n", err)
	}

	// Просматриваем данные проекта
	pr.PrintInfo()

	// Обновляем описание задачи 1
	if err := tk1.UpdateDescription("Новое описание важной задачи №2"); err != nil {
		log.Fatalf("task update description error: %v\n", err)
	}
	// Обновляем задачу в проекте
	if err := pr.UpdateTask(*tk1); err != nil {
		log.Fatalf("update task error: %v\n", err)
	}

	// Закрываем задачу 2
	if err := tk2.Close(); err != nil {
		log.Fatalf("task close error: %v\n", err)
	}
	// Обновляем задачу в проекте
	if err := pr.UpdateTask(*tk2); err != nil {
		log.Fatalf("update task error: %v\n", err)
	}

	// Просматриваем данные проекта
	pr.PrintInfo()

	// Отображаем только закрытые задачи
	fmt.Println("Закрытые задачи проекта:")
	for _, task := range pr.FilterTasksByStatus(project.StatusClosed) {
		fmt.Printf("Задача: %s, Описание: %s, Статус: %s\n", task.Title, task.Description, task.Status)
	}
	
	// Просматриваем данные проекта
	pr.PrintInfo()
}
