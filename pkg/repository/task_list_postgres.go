package repository

import (
	"fmt"
	"github.com/bwjson/toDoApp"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TaskListPostgres struct {
	db *sqlx.DB
}

func NewTaskListPostgres(db *sqlx.DB) *TaskListPostgres {
	return &TaskListPostgres{db: db}
}

func (t *TaskListPostgres) Create(userId int, list toDoApp.TaskList) (int, error) {
	tx, err := t.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", tasksListsTable)

	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		tx.Rollback()
	}

	return id, tx.Commit()
}

func (t *TaskListPostgres) GetAll(userId int) ([]toDoApp.TaskList, error) {
	var taskLists []toDoApp.TaskList

	getAllListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", tasksListsTable, usersListsTable)
	err := t.db.Select(&taskLists, getAllListsQuery, userId)

	return taskLists, err
}

func (t *TaskListPostgres) GetById(userId int, ListId int) (toDoApp.TaskList, error) {
	var taskList toDoApp.TaskList

	getListByIdQuery := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description 
											FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id
											WHERE ul.user_id = $1 AND ul.list_id = $2`, tasksListsTable, usersListsTable)
	err := t.db.Get(&taskList, getListByIdQuery, userId, ListId)

	return taskList, err
}

func (t *TaskListPostgres) Delete(userId int, ListId int) error {
	deleteListQuery := fmt.Sprintf(`DELETE FROM %s tl USING %s ul
 										   WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2`, tasksListsTable, usersListsTable)

	_, err := t.db.Exec(deleteListQuery, userId, ListId)

	return err
}

func (t *TaskListPostgres) Update(userId, ListId int, input toDoApp.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		tasksListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, ListId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := t.db.Exec(query, args...)
	return err
}
