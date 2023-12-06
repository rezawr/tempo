package handlers

import (
	"net/http"
	"tempo/helpers"
	services "tempo/services/news"

	"github.com/gin-gonic/gin"
)

type handlerShowNews struct {
	service services.ServiceShowNews
}

func NewHandlerShowNews(service services.ServiceShowNews) *handlerShowNews {
	return &handlerShowNews{service: service}
}

func (h *handlerShowNews) ShowNewsHandler(ctx *gin.Context) {
	ID := ctx.Param("id")
	if ID == "" {
		helpers.APIResponse(ctx, "unauthorized", http.StatusInternalServerError, http.MethodPut, nil)
		return
	}

	res, err := h.service.ShowNewsService(ID)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "News not found", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "", http.StatusOK, http.MethodPost, res)
	}
}
