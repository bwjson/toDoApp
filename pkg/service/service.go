package service

import (
	"github.com/bwjson/toDoApp"
	"github.com/bwjson/toDoApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user toDoApp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TaskList interface {
	Create(userId int, user toDoApp.TaskList) (int, error)
	GetAll(userId int) ([]toDoApp.TaskList, error)
	GetById(userId int, ListId int) (toDoApp.TaskList, error)
	Delete(userId int, ListId int) error
	Update(userId, ListId int, input toDoApp.UpdateListInput) error
}

type TaskItem interface {
	Create(userId, listId int, input toDoApp.TaskItem) (int, error)
	GetAll(userId, listId int) ([]toDoApp.TaskItem, error)
	GetById(userId, itemId int) (toDoApp.TaskItem, error)
	Delete(userId int, itemId int) error
	Update(userId, itemId int, input toDoApp.UpdateItemInput) error
}

type Service struct {
	Authorization
	TaskList
	TaskItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TaskList:      NewTaskListService(repos.TaskList),
		TaskItem:      NewTaskItemService(repos.TaskItem, repos.TaskList),
	}
}
