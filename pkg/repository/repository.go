package repository

type Authorization interface {
}

type TaskList interface {
}

type TaskItem interface {
}

type Repository struct {
	Authorization
	TaskList
	TaskItem
}

func NewRepository() *Repository {
	return &Repository{}
}
