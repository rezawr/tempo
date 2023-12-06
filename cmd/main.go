package main

import (
	"tempo/pkg"

	"github.com/sirupsen/logrus"
)

func main() {
	app := SetupRouter()
	logrus.Fatal(app.Run(":" + pkg.GodotEnv("GO_PORT")))
}
