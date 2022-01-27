package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	UUID      string `gorm:"primaryKey"`
	Email     string `gorm:"type:string;"`
	Password  string `gorm:"type:string;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Tokens struct {
	UUID         string `gorm:"primaryKey"`
	IdUser       string `gorm:"foreignkey:Users;references:UUID"`
	RefreshToken string `gorm:"type:string;"`
}

func ConnectionDatabase() (*gorm.DB, error) {
	fmt.Println("Connecting to Database ....")
	//db, err := sql.Open("mysql", "root:fanzru@tcp(103.55.38.98:1000)/petcare?charset=utf8mb4&parseTime=True")
	dsn := "root:fanzru@tcp(103.55.38.98:1000)/auth"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	// migrate to database
	err = db.AutoMigrate(&Users{})
	if err != nil {
		fmt.Println("Migration failed")
		return nil, err
	}
	// migrate to database
	err = db.AutoMigrate(&Tokens{})
	if err != nil {
		fmt.Println("Migration failed")
		return nil, err
	}

	fmt.Println("Database Connected")
	return db, nil
}
