package main

import (
	"fmt"
	"os"
	"tempo/models"
	"tempo/pkg"
	"tempo/routes"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	db := SetupDatabase()
	app := gin.Default()
	fmt.Println(db)

	if pkg.GodotEnv("GO_ENV") != "production" && pkg.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if pkg.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)

	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	app.Use(helmet.Default())
	app.Use(gzip.Gzip(gzip.BestCompression))

	routes.InitAuthRoutes(db, app)
	routes.InitNewsRoutes(db, app)

	return app
}

func SetupDatabase() *gorm.DB {
	fmt.Println(pkg.GodotEnv("DATABASE_URI"))
	db, err := gorm.Open(mysql.Open(pkg.GodotEnv("DATABASE_URI")), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connect into Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connect into Database Successfully")
	}

	err = db.AutoMigrate(
		&models.Users{},
		&models.News{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Database Migration Successfully")
	}

	return db
}
