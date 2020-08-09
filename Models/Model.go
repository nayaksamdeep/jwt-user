package Models

import (
//   "strings"
     "github.com/jinzhu/gorm"
   "time"
)

var DB *gorm.DB

type RedirectUrl struct {
	ID int64        `json:"id"`
        USERID int64     `json:"userid"`
	Url string       `form:"url" json:"url" binding:"required"`
	TinyUrl string `json:"tinyurl"`
	CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
	ExpiryDate time.Time `gorm:"default.CURRENT_TIMESTAMP" json:"expiry_date"`
}


/*
type User struct {
	ID        int64    `gorm:"primary_key;auto_increment" json:"id"`
	Name  	  string    `gorm:"size:255;not null;unique" json:"name"`
        Role	  bool      `json:"role"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
*/

// User is either created or retrieved and authentiacted user from google
type User struct {
	ID            int64    `gorm:"primary_key;auto_increment" json:"id"`
        Sub           string `json:"sub"`
        Name          string `json:"name"`
        GivenName     string `json:"given_name"`
        FamilyName    string `json:"family_name"`
        Profile       string `json:"profile"`
        Picture       string `json:"picture"`
        Email         string `json:"email"`
        EmailVerified bool   `json:"email_verified"`
        Gender        string `json:"gender"`
        GoogleUser	  bool      `json:"googleuser"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}


