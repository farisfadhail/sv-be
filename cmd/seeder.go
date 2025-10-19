package main

import (
	"fmt"
	"os"
	"strings"
	"test-be/config"
	"test-be/database/seeders"
	"test-be/utils"
)

// Code dari project sebelumnya
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run seeder.go [command]")
		return
	}

	command := os.Args[1]

	switch {
	case command == "seed:run":
		fmt.Println("Running seeders...")
		db := config.ConnectGormDB()
		seeders.RunAllSeeder(db)
		fmt.Println("Seeders done.")
	case command == "seed:create":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go seed:create <name>")
			return
		}

		createSeeder(os.Args[2])
	default:
		fmt.Println("Unknown command:", command)
		printSeederHelp()
	}
}

func createSeeder(name string) {
	if !strings.HasSuffix(name, "Seeder") {
		name += "Seeder"
	}

	funcName := name
	snakeFileName := utils.CamelToSnake(name) + ".go"
	dir := "database/seeders"
	filePath := dir + "/" + snakeFileName

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("ResponseFailed creating directory:", err)
			return
		}
	}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Seeder already exists:", filePath)
		return
	}

	content := fmt.Sprintf(`package seeders

import "gorm.io/gorm"

func %s(db *gorm.DB) {
    // TODO: implement %s
}
`, funcName, funcName)

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("ResponseFailed creating seeder:", err)
		return
	}

	fmt.Println("Seeder created at:", filePath)
}

func printSeederHelp() {
	fmt.Print(`
Usage:
  go run cmd/seeder.go <command> [options]

Available Commands:
  seed:run                Run all seeders defined in runSeeders()
  seed:create <Name>      AddNewData a new seeder file in database/seeders/
  --help                  Show this help message

Examples:
  go run main.go seed:create AdminRootSeeder
  go run main.go seed:run
`)
}
