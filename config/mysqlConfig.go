package config

import (
	"GoEchoAuth/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDatabase() (*gorm.DB, error) {
	fmt.Println("Connecting to Database ....")
	//db, err := sql.Open("mysql", "root:fanzru@tcp(103.55.38.98:1000)/petcare?charset=utf8mb4&parseTime=True")
	dsn := "affan:Affan!080701@tcp(103.55.38.98:3306)/auth"
	// dsn := "user:fanzru@tcp(103.55.38.98:3306)/auth"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	fmt.Println("Database Connected")
	return db, nil
}

func MigrationDB() (*gorm.DB, error) {
	// migrate to database
	db, err := ConnectionDatabase()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Users{})
	if err != nil {
		fmt.Println("Migration failed")
		return nil, err
	}
	// migrate to database
	err = db.AutoMigrate(&models.Tokens{})
	if err != nil {
		fmt.Println("Migration failed")
		return nil, err
	}
	return db, nil
}
