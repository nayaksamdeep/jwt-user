package Models

import (
	"fmt"
        "log"
//	"time"
        _"github.com/mattn/go-sqlite3"
//        "github.com/jinzhu/gorm"
)

func ListUsers(userstruct *[]User) (err error) {
	if err = DB.Find(userstruct).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func CreateUser(userstruct *User) (err error) {
	if err = DB.Create(userstruct).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetUser(userstruct *User, id string) (err error) {
	if err := DB.Where("id = ?", id).First(userstruct).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func UpdateUser(userstruct *User, id string) (err error) {
	fmt.Println(userstruct)
	DB.Save(userstruct)
	return nil
}

func DeleteUser(userstruct *User, id string) (err error) {
	DB.Where("id = ?", id).Delete(userstruct)
	return nil
}
