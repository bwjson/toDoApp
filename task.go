package toDoApp

type TaskList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TaskItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
