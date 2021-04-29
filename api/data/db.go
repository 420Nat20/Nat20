package data

import (
	"github.com/420Nat20/Nat20/nat-20/data/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB opens a database connection.
func NewDB() *gorm.DB {
	err := godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&model.GameModel{},
		&model.LocationModel{},
		&model.SubLocationModel{},
		&model.UserModel{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
