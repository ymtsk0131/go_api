package main

import (
	"go_api/config"
	"go_api/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Printf("test")
}
