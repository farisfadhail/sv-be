package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"test-be/config"
	"test-be/database/seeders"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Code dari project sebelumnya
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run migrate.go <command>")
	}

	command := os.Args[1]
	db := config.ConnectMigrationDB()

	switch command {
	case "migrate:up":
		runMigrateUp(db)
	case "migrate:create":
		runMigrateCreate()
	case "migrate:rollback":
		runMigrateRollback(db)
	case "migrate:fresh":
		runMigrateFresh(db)
	case "migrate:fresh-seed":
		runMigrateFreshAndSeed(db)
	case "migrate:version":
		runMigrateVersion(db)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printMigrateHelp()
		os.Exit(1)
	}
}

func runMigrateUp(db *migrate.Migrate) {
	err := db.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatalf("migration up failed: %s", err)
	}
	fmt.Println("Migration success!")
}

func runMigrateCreate() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run migrate.go migrate:create <migration_name>")
	}

	name := os.Args[2]

	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "database/migrations", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running migrate create: %v\n", err)
	}

	fmt.Println("Migration file created successfully.")
}

func runMigrateRollback(db *migrate.Migrate) {
	step := parseRollbackStep()
	err := db.Steps(-step)
	if err != nil {
		log.Fatalf("Error rollback migration: %v\n", err)
	}
	fmt.Printf("Rollback %d step(s) success!\n", step)
}

func parseRollbackStep() int {
	step := 1
	for _, arg := range os.Args[2:] {
		if len(arg) > 7 && arg[:7] == "--step=" {
			if val, err := strconv.Atoi(arg[7:]); err == nil {
				step = val
			}
		}
	}
	return step
}

func runMigrateFresh(db *migrate.Migrate) {
	if !confirmFresh() {
		fmt.Println("Migration fresh cancelled.")
	}

	version, dirty, err := db.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		log.Fatalf("Error get version: %v\n", err)
	}

	if dirty {
		fmt.Printf("Current version %d is dirty, forcing to set it clean.\n", version)
		if err := db.Force(int(version)); err != nil {
			log.Fatalf("Error force migration version: %v\n", err)
		}
		fmt.Println("Force version success!")
	}

	dropAllForeignKeys()

	if err := db.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Error running migration: %v\n", err)
	}
	fmt.Println("Rollback all migration success!")

	if err := db.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("Error migrate after fresh: %v\n", err)
	}
	fmt.Println("Fresh migration success!")
}

func runMigrateFreshAndSeed(db *migrate.Migrate) {
	runMigrateFresh(db)

	conn := config.ConnectGormDB()

	fmt.Println("Running seeders...")
	seeders.RunAllSeeder(conn)
	fmt.Println("Seeders done.")
}

func dropAllForeignKeys() {
	conn := config.ConnectGormDB()
	type Result struct {
		TableName      string
		ConstraintName string
	}

	var results []Result

	err := conn.Raw(`
		SELECT 
			tc.table_name, 
			tc.constraint_name
		FROM information_schema.table_constraints tc
		WHERE constraint_type = 'FOREIGN KEY'
	`).Scan(&results).Error

	if err != nil {
		log.Fatalf("Error getting foreign keys: %v\n", err)
	}

	for _, row := range results {
		sql := fmt.Sprintf("ALTER TABLE \"%s\" DROP CONSTRAINT \"%s\" CASCADE;", row.TableName, row.ConstraintName)
		if err := conn.Exec(sql).Error; err != nil {
			log.Fatalf("Error dropping foreign key %s on table %s: %v\n", row.ConstraintName, row.TableName, err)
		}
		fmt.Printf("Dropped foreign key %s on table %s\n", row.ConstraintName, row.TableName)
	}
}

func confirmFresh() bool {
	var confirm string
	fmt.Print("Are you sure you want to run fresh migration? This will drop all tables. (y/N): ")
	_, _ = fmt.Scanln(&confirm)
	return confirm == "y"
}

func runMigrateVersion(db *migrate.Migrate) {
	version, dirty, err := db.Version()
	if err != nil {
		log.Fatalf("Error getting migration version: %v\n", err)
	}
	fmt.Printf("Current migration version: %d, dirty: %v\n", version, dirty)
}

func printMigrateHelp() {
	fmt.Print(`
Usage:
  go run migrate.go <command> [options]

Available Commands:
  migrate:up              Apply all up migrations
  migrate:rollback        Rollback last migration step (use --step=N to rollback multiple steps)
  migrate:create <name>   AddNewData new migration file with given name
  migrate:fresh           Drop all tables and re-run all migrations (not allowed in production)
  migrate:fresh-seed      Drop all tables and re-run all migrations and seeding (not allowed in production)
  migrate:version         Show current migration version

Options:
  --step=N                Number of steps to rollback (default: 1)

Examples:
  go run migrate.go migrate:up
  go run migrate.go migrate:rollback --step=2
  go run migrate.go migrate:create create_users_table
  go run migrate.go migrate:fresh
  go run migrate.go migrate:fresh-seed
  go run migrate.go migrate:version
`)
}
