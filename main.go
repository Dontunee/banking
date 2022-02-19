package main

import (
	"github.com/Dontunee/banking/app"
	"github.com/Dontunee/banking/logger"
)

func main() {
	logger.Info("Application is about to start....")
	app.Start()
}
