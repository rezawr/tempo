package handlers

import (
	"net/http"
	"tempo/helpers"
	services "tempo/services/news"

	"github.com/gin-gonic/gin"
)

type handlerGetNews struct {
	service services.ServiceGetNews
}

func NewHandlerGetNews(service services.ServiceGetNews) *handlerGetNews {
	return &handlerGetNews{service: service}
}

func (h *handlerGetNews) GetNewsHandler(ctx *gin.Context) {
	res, err := h.service.GetNewsService()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "News cant be retrieve", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "", http.StatusOK, http.MethodPost, res)
	}
}
