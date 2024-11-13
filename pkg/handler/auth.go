package handler

import (
	"github.com/bwjson/toDoApp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Получаем и валидируем данные пользователя через JSON
func (h *Handler) signUp(c *gin.Context) {
	var input toDoApp.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
