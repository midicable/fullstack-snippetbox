package handler

import (
	"fmt"
	"fullstack-snippetbox-backend/internal/snippets/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SnippetsHandler struct {
	SnippetsService *service.SnippetsService
}

func NewSnippetsHandler(snippetsService *service.SnippetsService) *SnippetsHandler {
	return &SnippetsHandler{
		SnippetsService: snippetsService,
	}
}

func (h *SnippetsHandler) GetSnippet(c *gin.Context) {
	id := c.Param("id")
	res := h.SnippetsService.GetSnippet()

	fmt.Println(res)

	c.JSON(http.StatusOK, map[string]any{
		"data": res,
		"meta": fmt.Sprintf("Успешно получен snippet с id = %s", id),
	})
}