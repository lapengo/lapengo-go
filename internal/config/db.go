package config

import (
	"fmt"
	"log"
	"os"

	"github.com/lapengo/lapengo-go/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var env string
	if Config("PROD") == "true" {
		env = "PROD"
	} else {
		env = "DEV"
	}

	DBHost := Config("DB_HOST__" + env)
	DBUser := Config("DB_USER__" + env)
	DBPassword := Config("DB_PASSWORD__" + env)
	DBPort := Config("DB_PORT__" + env)
	DBName := Config("DB_NAME__" + env)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)

	DBLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{Logger: DBLogger})
	// DB, err := gorm.Open(postgres.Open(dbURL))

	helper.PanicIfError(err)
	// fmt.Println("Database is successfully connected")

	DBConn = DB
}
