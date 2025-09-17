package main

import (
	"log"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/models"
)

func init() {
	conf.LoadEnvVars()
	conf.ConnectToDatabase()
}

func main() {
	log.Println("Starting migration...")

	conf.DB.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Team{},
		&models.Bug{},
	)

	log.Println("Migration completed successfully.")
}
