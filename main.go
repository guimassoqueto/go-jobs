package main

import (
	"github.com/guimassoqueto/go-jobs/config"
	"github.com/guimassoqueto/go-jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main") 
	// Init config
	err := config.Init()
	if err != nil {
		logger.Error("Config initialization failed", err)
		return
	}

	router.Initialize()
}