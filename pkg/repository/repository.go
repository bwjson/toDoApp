package repository

import (
	"github.com/bwjson/toDoApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user toDoApp.User) (int, error)
	GetUser(username, password string) (toDoApp.User, error)
}

type TaskList interface {
	Create(userId int, list toDoApp.TaskList) (int, error)
	GetAll(userId int) ([]toDoApp.TaskList, error)
	GetById(userId int, ListId int) (toDoApp.TaskList, error)
	Delete(userId int, ListId int) error
	Update(userId, ListId int, input toDoApp.UpdateListInput) error
}

type TaskItem interface {
}

type Repository struct {
	Authorization
	TaskList
	TaskItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TaskList:      NewTaskListPostgres(db),
	}
}
