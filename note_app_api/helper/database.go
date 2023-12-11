package helper

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ConnectionOptions struct {
	dialect  string
	host     string
	port     string
	username string
	password string
	database string
}

// Method to connect to databases like MySQL, PostgreSQL, SQLite, etc.
// 
// This method will return a pointer to gorm.DB
func Connect(opts ConnectionOptions) *gorm.DB {
	var db *gorm.DB

	switch opts.dialect {
	case "sqlite":
		conn, err := gorm.Open(sqlite.Open(opts.database), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		db = conn
	default:
		panic("unknown database dialect")
	}

	return db
}
