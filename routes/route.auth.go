package routes

import (
	handlers "tempo/handlers/auth"
	repositories "tempo/repositories/auth"
	services "tempo/services/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	/**
	| ----------------------------------------------------
	| Describe all handler, services, and repositories
	| ----------------------------------------------------
	*/

	registerRepository := repositories.NewRepositoryRegister(db)
	registerService := services.NewServiceRegister(registerRepository)
	registerHandler := handlers.NewHandlerRegister(registerService)

	loginRepository := repositories.NewRepositoryLogin(db)
	loginService := services.NewServiceLogin(loginRepository)
	loginHandler := handlers.NewHandlerLogin(loginService)

	/**
	| ----------------------------------------------------
	| Describe all routes
	| ----------------------------------------------------
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
