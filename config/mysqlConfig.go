package config

import (
	"GoEchoAuth/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDatabase() (*gorm.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/auth?charset=utf8&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrationDB() (*gorm.DB, error) {
	// migrate to database
	fmt.Println("Connecting to Database ....")
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
	fmt.Println("Database Connected")
	return db, nil
}
