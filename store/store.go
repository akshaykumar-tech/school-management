package store

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (store *Postgres) NewStore() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_school port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	} else {
		store.DB = db
	}

	fmt.Println("✅ Connected to PostgreSQL via GORM")

}


type StoreOperation interface {

	NewStore()

}
