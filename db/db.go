package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	CONNECTION_PROD           = "production"
	CONNECTION_DEV            = "development"
	CONNECTION_DEV_PERSISTENT = "dev-persistent"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Init() {
	_db, err := gorm.Open(
		getDBConnection(os.Getenv("DB_STATE")),
		&gorm.Config{TranslateError: true},
	)
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	err = _db.AutoMigrate(
		&entity.Payment{},
		&entity.WashingMachine{},
		&entity.Service{},
		&entity.Order{},
		&entity.Partner{},
		&entity.Customer{},
	)
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	db = _db
}

func GetConnection() *gorm.DB {
	once.Do(Init)
	return db
}

func getDBConnection(connection string) gorm.Dialector {
	switch connection {
	case CONNECTION_DEV:
		return sqlite.Open("file::memory:?cache=shared")
	case CONNECTION_DEV_PERSISTENT:
		return sqlite.Open("data.db")
	case CONNECTION_PROD:
		return postgres.Open(buildConnectionString())
	}

	panic(fmt.Sprintf("Invalid database connection: %s", connection))
}

func buildConnectionString() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port)
}

func MigrateDummy() {
	insertDataDummy()
}
