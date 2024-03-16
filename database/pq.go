package database

import (
	"fmt"
	"golangproject/models/project"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

//make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {

	// load .env file
	godotenv.Load(".env")
	var err error
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")
   
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
   
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
   
	if err != nil {
	 panic(err)
	} else {
	 fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}
}

func AutoMigrateModels() {

	Database.AutoMigrate(&project.Project{})
}