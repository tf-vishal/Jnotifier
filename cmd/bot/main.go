package main

import (
	"log"

	"github.com/tf-vishal/JNotifier/internal/schedulers"
	"github.com/tf-vishal/JNotifier/internal/server"
	"github.com/tf-vishal/JNotifier/internal/utils"
)

func main() {
	utils.LoadEnv()
	log.Println("Starting JNotifier Bot...")
	go server.StartWebServer()
	go schedulers.StartJobScheduler()
	select {}
}
