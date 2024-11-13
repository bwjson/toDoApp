package repository

import (
	"github.com/bwjson/toDoApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user toDoApp.User) (int, error)
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
