package config

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConfigurationDB() (host, port, user, password, dbname string) {
	LoadEnv()

	host = GetEnv("DB_HOST", "localhost")
	port = GetEnv("DB_PORT", "3306")
	user = GetEnv("DB_USER", "faris")
	password = GetEnv("DB_PASSWORD", "password")
	dbname = GetEnv("DB_NAME", "bri-edc")

	return
}

func ConnectGormDB() *gorm.DB {
	host, port, user, password, dbname := ConfigurationDB()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect to database")
	}

	appEnv := GetEnv("APP_ENV", "local")

	if appEnv == "local" {
		//db = db.Debug()
	}

	return db
}

func ConnectMigrationDB() *migrate.Migrate {
	host, port, user, password, dbname := ConfigurationDB()

	dbURL := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname,
	)

	m, err := migrate.New(
		"file://database/migrations",
		dbURL,
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database migration: %v", err))
	}

	return m
}
