package main

import (
	"context"
	"log"
	"os"

	"github.com/badtripdude/gh-analyzer/internal/app"
	"github.com/badtripdude/gh-analyzer/internal/db"
	"github.com/badtripdude/gh-analyzer/internal/kafka"
)

// ! main initializes the database connection, runs migrations, and starts the Kafka consumer.
// ? this is a basic version; further expansions are planned.
func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	if err := db.RunMigrations(database); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	service := app.NewService(database)

	consumer := kafka.NewConsumer(
		[]string{os.Getenv("KAFKA_BROKER")},
		"analysis.results",
		"storage-group",
		service,
	)

	log.Println("Starting consumer...")
	if err := consumer.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
