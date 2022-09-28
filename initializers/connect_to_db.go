package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() {
	var err error
	// postgres://ndiqionr:4BQHurAUBN_GUb0II1RDLaxk-JEl5KpM@babar.db.elephantsql.com/ndiqionr
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(DB)
	if err != nil {
		panic("Failed to connect database")
	}
}
