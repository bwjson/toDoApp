package service

import (
	"github.com/bwjson/toDoApp"
	"github.com/bwjson/toDoApp/pkg/repository"
)

type taskItemService struct {
	repo     repository.TaskItem
	listRepo repository.TaskList
}

func NewTaskItemService(repo repository.TaskItem, listRepo repository.TaskList) *taskItemService {
	return &taskItemService{repo: repo, listRepo: listRepo}
}

func (t *taskItemService) Create(userId, listId int, input toDoApp.TaskItem) (int, error) {
	_, err := t.listRepo.GetById(userId, listId)

	// list does not exist
	if err != nil {
		return 0, err
	}

	return t.repo.Create(listId, input)
}

func (t *taskItemService) GetAll(userId, listId int) ([]toDoApp.TaskItem, error) {
	return t.repo.GetAll(userId, listId)
}

func (t *taskItemService) GetById(userId, itemId int) (toDoApp.TaskItem, error) {
	return t.repo.GetById(userId, itemId)
}

func (t *taskItemService) Delete(userId, itemId int) error {
	return t.repo.Delete(userId, itemId)
}

func (t *taskItemService) Update(userId, itemId int, input toDoApp.UpdateItemInput) error {
	return t.repo.Update(userId, itemId, input)
}
