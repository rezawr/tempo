package handlers

import (
	"net/http"
	"tempo/helpers"
	"tempo/pkg"
	serializers "tempo/serializers/news"
	services "tempo/services/news"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handlerUpdateNews struct {
	service services.ServiceUpdateNews
}

func NewHandlerUpdateNews(service services.ServiceUpdateNews) *handlerUpdateNews {
	return &handlerUpdateNews{service: service}
}

func (h *handlerUpdateNews) UpdateNewsHandler(ctx *gin.Context) {

	var input serializers.NewsSerializer
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Title",
				Message: "title is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Body",
				Message: "body is required on body",
			},
		},
	}

	errResponse, errCount := pkg.GoValidator(&input, config.Options)
	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPut, errResponse)
		return
	}

	uID, ok := ctx.Get("user_id")
	if !ok {
		helpers.APIResponse(ctx, "unauthorized", http.StatusUnauthorized, http.MethodPut, nil)
		return

	}
	userID := uID.(string)

	ID := ctx.Param("id")
	if ID == "" {
		helpers.APIResponse(ctx, "unauthorized", http.StatusInternalServerError, http.MethodPut, nil)
		return
	}

	res, err := h.service.UpdateNewsService(&input, ID, userID)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "News cant be updated", err.Code, http.MethodPut, nil)
		return
	default:
		helpers.APIResponse(ctx, "News successfully Updated", http.StatusOK, http.MethodPut, res)
	}
}
