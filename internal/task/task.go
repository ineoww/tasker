package task

import "time"


type Task struct {
	Id int	`json:"id"`
	Description string	`json:"description"`
	Status Status	`json:"status"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type Status string 

	const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

