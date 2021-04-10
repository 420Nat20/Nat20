package data

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open a database connection.
func NewDB() *gorm.DB {
	err := godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&GameModel{},
		&LocationModel{},
		&SubLocationModel{},
		&UserModel{},
	)

	return db
}
