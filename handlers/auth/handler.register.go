package handlers

import (
	"net/http"
	"tempo/helpers"
	"tempo/pkg"
	serializers "tempo/serializers/auth"
	services "tempo/services/auth"

	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/gin-gonic/gin"
)

type handlerRegister struct {
	service services.ServiceRegister
}

func NewHandlerRegister(service services.ServiceRegister) *handlerRegister {
	return &handlerRegister{service: service}
}

func (h *handlerRegister) RegisterHandler(ctx *gin.Context) {

	var input serializers.RegisterSerializer
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Fullname",
				Message: "fullname is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "email",
				Field:   "Email",
				Message: "email format is not valid",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Password",
				Message: "password is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "gte",
				Field:   "Password",
				Message: "password minimum must be 8 character",
			},
		},
	}

	errResponse, errCount := pkg.GoValidator(input, config.Options)

	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, err := h.service.RegisterService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "Email already exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Register new account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
		return
	}
}
