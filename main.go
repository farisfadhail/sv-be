package main

import (
	"fmt"
	"log"
	"test-be/config"
	"test-be/internal/injector"
	"test-be/internal/routes"
)

func main() {
	ct, err := injector.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	app := config.NewFiber()

	routes.SetupRouter(app, ct)

	port := config.GetEnv("PORT", "8080")
	url := config.GetEnv("URL", "0.0.0.0")
	err = app.Listen(fmt.Sprintf("%s:%s", url, port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}
