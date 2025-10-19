package seeders

import (
	"test-be/config"

	"gorm.io/gorm"
)

func RunAllSeeder(db *gorm.DB) {
	config.LoadEnv()
}
