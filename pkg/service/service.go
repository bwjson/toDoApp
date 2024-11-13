package service

import (
	"github.com/bwjson/toDoApp"
	"github.com/bwjson/toDoApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user toDoApp.User) (int, error)
}

type TaskList interface {
}

type TaskItem interface {
}

type Service struct {
	Authorization
	TaskList
	TaskItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
