package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Open a database connection.
func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&UserModel{},
	)

	return db
}
