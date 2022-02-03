package models

import "time"

type Users struct {
	UUID      int64     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Email     string    `gorm:"type:string;unique"`
	Password  string    `gorm:"type:string;"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;"`
}
