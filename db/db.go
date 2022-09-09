package db

import (
	_ "database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reviewer-api-service/model"
	"sync"
	"time"
)

var txdbInitialized bool
var mutex sync.Mutex

func dsn() (string, error) {
	err := godotenv.Load("env/test.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.New("$DB_HOST is not set")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return "", errors.New("$DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return "", errors.New("$DB_PASSWORD is not set")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return "", errors.New("$DB_NAME is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.New("$DB_PORT is not set")
	}

	//"user:password@host:port/dbname?option1&option2"
	return fmt.Sprintf("%s:%s@(%s:%s)/%s",
		user, password, host, port, name), nil
}
func New() (*gorm.DB, error) {
	s, err := dsn()
	if err != nil {
		return nil, err
	}

	var d *gorm.DB
	for i := 0; i < 10; i++ {
		d, err = gorm.Open("mysql", s)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return nil, err
	}
	fmt.Printf("Connect success " + s)

	d.DB().SetMaxIdleConns(3)
	d.LogMode(false)

	return d, nil
}

// AutoMigrate is a wrapper of (*gorm.DB).AutoMigrate
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.UserRequests{},
	).Error
	if err != nil {
		return err
	}
	return nil
}
