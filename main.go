/*
 * This program measures the number of glasses of water a person drinks a day
 * It will an user to input the number of glasses he has drank
 * It will also show him the total number of cups he has consumed for the day
 */

package main

import  (
     "fmt"
     "log"
     "github.com/nayaksamdeep/jwt-user/Models"
     "github.com/nayaksamdeep/jwt-user/Routes"
     "github.com/jinzhu/gorm"
     _"github.com/mattn/go-sqlite3"
)

var err error

func main() {

	// Creating a connection to the database

	Models.DB, err = gorm.Open("sqlite3", "./tinyurlgorm.db")

	if err != nil {
                log.Fatal(err)
	}

	defer Models.DB.Close()

	// run the migrations: todo struct
	Models.DB.AutoMigrate(&Models.RedirectUrl{})

        fmt.Println("hi there! Welcome to Contact List Service\n");

        r := Routes.SetupRouter()
        // running
        r.Run() 
}
