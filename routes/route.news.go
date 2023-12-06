package routes

import (
	handlers "tempo/handlers/news"
	middleware "tempo/middlewares"
	repositories "tempo/repositories/news"
	services "tempo/services/news"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitNewsRoutes(db *gorm.DB, route *gin.Engine) {
	/**
	| ----------------------------------------------------
	| Describe all handler, services, and repositories
	| ----------------------------------------------------
	*/
	getRepository := repositories.NewRepositoryGetNews(db)
	getService := services.NewServiceGetNews(getRepository)
	getHandler := handlers.NewHandlerGetNews(getService)

	createRepository := repositories.NewRepositoryCreateNews(db)
	createService := services.NewServiceCreateNews(createRepository)
	createHandler := handlers.NewHandlerCreateNews(createService)

	showRepository := repositories.NewRepositoryShowNews(db)
	showService := services.NewServiceShowNews(showRepository)
	showHandler := handlers.NewHandlerShowNews(showService)

	updateRepository := repositories.NewRepositoryUpdateNews(db)
	updateService := services.NewServiceUpdateNews(updateRepository)
	updateHandler := handlers.NewHandlerUpdateNews(updateService)

	deleteRepository := repositories.NewRepositoryDeleteNews(db)
	deleteService := services.NewServiceDeleteNews(deleteRepository)
	deleteHandler := handlers.NewHandlerDeleteNews(deleteService)

	/**
	| ----------------------------------------------------
	| Describe all routes
	| ----------------------------------------------------
	*/
	groupRoute := route.Group("/api/v1/news").Use(middleware.Auth())
	groupRoute.GET("/", getHandler.GetNewsHandler)
	groupRoute.GET("/:id", showHandler.ShowNewsHandler)
	groupRoute.POST("/", createHandler.CreateNewsHandler)
	groupRoute.PUT("/:id", updateHandler.UpdateNewsHandler)
	groupRoute.DELETE("/:id", deleteHandler.DeleteNewsHandler)
}
