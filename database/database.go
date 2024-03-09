package database

import (
	"fmt"
	"log"
	"os"

	"github.com/itoodua12/Fiber-Rest-API/model"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var Database DbInstance

func ConnectDB() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	
	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	log.Println("Running Migrations")

	db.Logger = logger.Default.LogMode(logger.Info)
	if err := db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}); err != nil {
		log.Fatalf("Failed to auto migrate tables: %v", err)
	}

	Database = DbInstance{DB: db}

}
