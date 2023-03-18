package database
import (
	"jwt-authentication-golang/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Instance *gorm.DB //Here, we are defining an instance of the database. This variable will be used across the entire application to communicate with the database.
var dbError error

// The Connect() function takes in the MySQL connection string (which we are going to pass from the main method shortly) and tries to connect to the database using GORM.
func Connect(connectionString string) () {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

// we will call this Migrate() function to ensure that in our database, there is a users table. If not present, GORM will automatically create a new table named “users” for us.
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}