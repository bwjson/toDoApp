package service

import (
	"github.com/bwjson/toDoApp"
	"github.com/bwjson/toDoApp/pkg/repository"
)

type taskListService struct {
	repo repository.TaskList
}

func NewTaskListService(repo repository.TaskList) *taskListService {
	return &taskListService{repo: repo}
}

func (s *taskListService) Create(userId int, list toDoApp.TaskList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *taskListService) GetAll(userId int) ([]toDoApp.TaskList, error) {
	return s.repo.GetAll(userId)
}

func (s *taskListService) GetById(userId int, ListId int) (toDoApp.TaskList, error) {
	return s.repo.GetById(userId, ListId)
}

func (s *taskListService) Delete(userId int, ListId int) error {
	return s.repo.Delete(userId, ListId)
}

func (s *taskListService) Update(userId, ListId int, input toDoApp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, ListId, input)
}
