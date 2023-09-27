package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic("failed to connect to database")
	}

	seedDB(db)
	return db
}
