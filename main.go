package main

import (
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

	err = app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}
