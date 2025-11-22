package sqlconnect

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	fmt.Println("Trying to connect to database...")

	connectionString := os.Getenv("CONNECTION_STRING")

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MariaDB: %w", err)
	}

	fmt.Println("Connected to MariaDB")
	return db, nil
}
