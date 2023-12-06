package handlers

import (
	"net/http"
	"tempo/helpers"
	services "tempo/services/news"

	"github.com/gin-gonic/gin"
)

type handlerDeleteNews struct {
	service services.ServiceDeleteNews
}

func NewHandlerDeleteNews(service services.ServiceDeleteNews) *handlerDeleteNews {
	return &handlerDeleteNews{service: service}
}

func (h *handlerDeleteNews) DeleteNewsHandler(ctx *gin.Context) {
	ID := ctx.Param("id")
	if ID == "" {
		helpers.APIResponse(ctx, "unauthorized", http.StatusInternalServerError, http.MethodPut, nil)
		return
	}

	_, err := h.service.DeleteNewsService(ID)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "news data is not exist or deleted", err.Code, http.MethodDelete, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Delete news data failed", err.Code, http.MethodDelete, nil)
		return
	default:
		helpers.APIResponse(ctx, "delete successful", http.StatusOK, http.MethodPost, nil)
	}
}
