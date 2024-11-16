package repository

import (
	"fmt"
	"github.com/bwjson/toDoApp"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TaskItemPostgres struct {
	db *sqlx.DB
}

func NewTaskItemPostgres(db *sqlx.DB) *TaskItemPostgres {
	return &TaskItemPostgres{db: db}
}

func (t *TaskItemPostgres) Create(listId int, item toDoApp.TaskItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (t *TaskItemPostgres) GetAll(userId, listId int) ([]toDoApp.TaskItem, error) {
	var items []toDoApp.TaskItem
	getAllItemsQuery := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := t.db.Select(&items, getAllItemsQuery, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (t *TaskItemPostgres) GetById(userId, itemId int) (toDoApp.TaskItem, error) {
	var item toDoApp.TaskItem

	getItemByIdQuery := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := t.db.Get(&item, getItemByIdQuery, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (t *TaskItemPostgres) Delete(userId, itemId int) error {
	deleteListQuery := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
 										   WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	_, err := t.db.Exec(deleteListQuery, userId, itemId)

	return err
}

func (t *TaskItemPostgres) Update(userId, itemId int, input toDoApp.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
								 WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := t.db.Exec(query, args...)
	return err
}
