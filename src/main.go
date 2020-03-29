package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sofyan48/nemo/src/config"

	"github.com/joho/godotenv"
	apiRouter "github.com/sofyan48/nemo/src/router"
	"github.com/sofyan48/nemo/src/worker"
)

// ConfigEnvironment ...
func ConfigEnvironment(env string) {
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	ConfigEnvironment(*environment)
	go worker.LoadWorker()
	startApp()
}

func startApp() {
	router := config.SetupRouter()
	apiRouter.LoadRouter(router)
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.Run(serverString)
}
