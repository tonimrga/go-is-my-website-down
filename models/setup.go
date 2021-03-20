// models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// ConnectDataBase - connection to the sqlite db
func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "websites.db")

	if err != nil {
		panic("Connection to the database failed!")
	}

	database.AutoMigrate(&Website{})

	DB = database
}
