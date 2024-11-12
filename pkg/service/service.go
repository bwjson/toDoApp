package service

import "github.com/bwjson/toDoApp/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
