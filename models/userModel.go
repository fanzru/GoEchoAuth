package models

import "time"

type Users struct {
	UUID      int64  `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Email     string `gorm:"type:string;"`
	Password  string `gorm:"type:string;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
