package Models

import (
	"fmt"
        "log"
	_"github.com/mattn/go-sqlite3"
//        "github.com/jinzhu/gorm"
)

//var DB *gorm.DB

func GetAllUrl(urlstruct  *[]RedirectUrl) (err error) {
	if err = DB.Find(urlstruct).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func ConvertAUrl(urlstruct *RedirectUrl) (err error) {

	if err = DB.Create(urlstruct).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func RedirectAUrl(urlstruct *RedirectUrl, id string) (err error) {
	if err := DB.Where("id = ?", id).First(urlstruct).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetUrl(urlstruct *RedirectUrl, id string) (err error) {
        if err := DB.Where("id = ?", id).First(urlstruct).Error; err != nil {
                log.Fatal(err)
                return err
        }
        return nil
}

func UpdateAUrl(urlstruct *RedirectUrl, id string) (err error) {
	fmt.Println(urlstruct)
	DB.Save(urlstruct)
	return nil

}

func DeleteAUrl(urlstruct *RedirectUrl, id string) (err error) {
	DB.Where("id = ?", id).Delete(urlstruct)
	return nil
}
