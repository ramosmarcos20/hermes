package migrations

import (
	"hermes/config"
	"hermes/internal/models/entities"
	"log"
)

func RunMigration() {
	env := config.LoadEnv()
	db, err := config.Connection(env)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatalf("Could not migrate models: %v", err)
	}
	log.Println("Migration completed successfully")
}
