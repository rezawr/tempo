package handlers

import (
	"net/http"
	"tempo/helpers"
	"tempo/pkg"
	serializers "tempo/serializers/news"
	services "tempo/services/news"

	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/gin-gonic/gin"
)

type handlerCreateNews struct {
	service services.ServiceCreateNews
}

func NewHandlerCreateNews(service services.ServiceCreateNews) *handlerCreateNews {
	return &handlerCreateNews{service: service}
}

func (h *handlerCreateNews) CreateNewsHandler(ctx *gin.Context) {

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
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	uID, ok := ctx.Get("user_id")
	if !ok {
		helpers.APIResponse(ctx, "unauthorized", http.StatusUnauthorized, http.MethodPost, nil)
		return

	}
	userID := uID.(string)

	res, err := h.service.CreateNewsService(&input, userID)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "News cant be stored", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "News successfully created", http.StatusOK, http.MethodPost, res)
	}
}
