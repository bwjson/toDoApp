package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
	SSLMode  string
}

const (
	usersTable      = "public.users"
	todoListsTable  = "public.todo_lists"
	usersListsTable = "public.users_lists"
	todoItemsTable  = "public.todo_items"
	listsItemsTable = "public.lists_items"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
