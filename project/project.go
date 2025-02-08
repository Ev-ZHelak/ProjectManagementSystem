package project

import (
	"fmt"
	"slices"
	"strings"

	"github.com/google/uuid"
)

const (
	StatusClosed = "Closed"
	StatusOpen   = "Open"
)

type Status string

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      Status
}

type Project struct {
	ID        uuid.UUID
	Name      string
	TasksList []Task
}

func New(u uuid.UUID, name string) (*Project, error) {
	if u == uuid.Nil {
		return nil, fmt.Errorf("invalid UUID")
	}

	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	return &Project{
		ID:   u,
		Name: name,
	}, nil
}

func NewTask(u uuid.UUID, title string, description string) (*Task, error) {
	if u == uuid.Nil {
		return nil, fmt.Errorf("invalid UUID")
	}

	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	if description == "" {
		return nil, fmt.Errorf("description cannot be empty")
	}

	return &Task{
		ID:          u,
		Title:       title,
		Description: description,
		Status:      StatusOpen,
	}, nil
}

func (p *Project) AddTask(tk Task) error {
	for _, t := range p.TasksList {
		if t.ID == tk.ID {
			return fmt.Errorf("id already exists")
		}
	}
	p.TasksList = append(p.TasksList, tk)
	return nil
}

func (p *Project) UpdateTask(tk Task) error {
	for i := 0; i < len(p.TasksList); i++ {
		if p.TasksList[i].ID == tk.ID {
			p.TasksList[i] = tk
			return nil
		}
	}
	return fmt.Errorf("id not exists")
}

func (p *Task) UpdateDescription(description string) error {
	if description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	if p.Status == StatusClosed {
		return fmt.Errorf("closed task")
	}
	p.Description = description
	return nil
}

func (p *Task) Close() error {
	if p.Status == StatusOpen {
		p.Status = StatusClosed
		return nil
	}
	return fmt.Errorf("closed task")
}

func (p Project) FilterTasksByStatus(status Status) []Task {
	result := slices.DeleteFunc(slices.Clone(p.TasksList), func(a Task) bool {
		return !(a.Status == status)
	})
	return result
}

func (p Project) PrintInfo() {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("Project Info\nID: %s\nName: %s\n", p.ID, p.Name)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("Tasks: ")
	for i, task := range p.TasksList {
		fmt.Println(strings.Repeat("-", 80))
		fmt.Printf("Task: %d\n", i+1)
		fmt.Printf(
			"ID: %s\nTitle: %s\nDescription: %s\nStatus: %s\n",
			task.ID, task.Title, task.Description,
			task.Status,
		)
	}
	fmt.Println(strings.Repeat("=", 80))
	fmt.Print("\n\n")
}
