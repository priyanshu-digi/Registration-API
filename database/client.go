package database

import (
	"jwt-authentication-golang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB //instance of DB
var dbError error

func Connect(connectionString string) {
	//The Connect() function takes in the MySQL connection string
	//and tries to connect to the database using GORM.
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	//Once connected to the database using the previous Connect() function,
	// we will call this Migrate() function to ensure that in our database,
	//there is a users table. If not present, GORM will automatically create
	// a new table named “users” for us.
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
