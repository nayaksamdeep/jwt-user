package Models

import (
//   "strings"
     "github.com/jinzhu/gorm"
   "time"
)

var DB *gorm.DB

type RedirectUrl struct {
	ID uint            `json:"id"`
	Url string       `form:"url" json:"url" binding:"required"`
	TinyUrl string `json:"tinyurl"`
	CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
	ExpiryDate time.Time `gorm:"default.CURRENT_TIMESTAMP" json:"expiry_date"`
}


type User struct {
	ID        int64    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Admin struct {
	ID int64            `json:"id"`
	Name string `json:"name"`
	Password  string    `json:"password"`
}
