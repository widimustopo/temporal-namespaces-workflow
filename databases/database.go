package database

import (
	"fmt"

	"github.com/widimustopo/temporal-namespaces-workflow/libs"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(config *libs.Config) *gorm.DB {
	dbHost := config.DatabaseHost
	dbPort := config.DatabasePort
	dbUser := config.DatabaseUser
	dbPass := config.DatabasePassword
	dbName := config.DatabaseName

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Errorf("failed to open database: %v", err)
	}

	return db
}
