package config

import (
	"os"

	"github.com/guimassoqueto/go-jobs/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



func InitializeSQLlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// check whether database file exists (only for sqllite)
	err := checkForSQLite(dbPath)
	if err != nil {
		logger.Error(err)
	}

	// create database connection 
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("error while stablishing database connection: %v", err)
		return nil, err
	}

	// database migration
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqllite automigration error: %v", err)
		return nil, err
	}
	
	return db, nil
}

// helper function to validate if the sqllite db exists
func checkForSQLite(dbPath string) error {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database not found, creating...")
		// create database file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}