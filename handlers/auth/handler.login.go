package handlers

import (
	"net/http"
	"tempo/helpers"
	"tempo/pkg"
	serializers "tempo/serializers/auth"
	services "tempo/services/auth"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handlerLogin struct {
	service services.ServiceLogin
}

func NewHandlerLogin(service services.ServiceLogin) *handlerLogin {
	return &handlerLogin{service: service}
}

func (h *handlerLogin) LoginHandler(ctx *gin.Context) {

	var input serializers.LoginSerializer
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
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
		},
	}

	errResponse, errCount := pkg.GoValidator(&input, config.Options)

	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	res, err := h.service.LoginService(&input)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(ctx, "User account is not registered", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(ctx, "Username or password is wrong", err.Code, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": res.ID, "email": res.Email}
		accessToken, errToken := pkg.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			helpers.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		helpers.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	}
}
