package database

import (
	"fmt"

	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/config"
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/model"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	// database info
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBPassword)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	err := db.AutoMigrate(
		model.Admin{},
	)
	if err != nil {
		return nil, err
	}

	return db, dbErr
}
