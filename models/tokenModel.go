package models

type Tokens struct {
	UUID         int64  `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	IdUser       string `gorm:"foreignkey:Users;references:UUID"`
	RefreshToken string `gorm:"type:string;"`
}
