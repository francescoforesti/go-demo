package main

import (
	"fmt"
	"github.com/francescoforesti/go-demo/gorm/models"
	"github.com/francescoforesti/go-demo/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logging.InitializeLoggers()

	dsn := "host=0.0.0.0 user=postgres password=postgres DB.name=postgres port=5432 sslmode=disable TimeZone=Europe/Rome"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.GormMessage{})

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&models.GormMessage{Message: "42", Reversed: "24"}).Error; err != nil {
			return err
		}
		if err := db.Create(&models.GormMessage{Message: "420", Reversed: "024"}).Error; err != nil {
			return err
		}
		return nil
	})

	var found []models.GormMessage
	var _ = db.Find(&found)
	logging.Info(fmt.Sprintf("%+v", found))

	db.Transaction(func(tx *gorm.DB) error {
		tx.Where("1 = 1").Delete(models.GormMessage{})
		return nil
	})

	var _ = db.Find(&found)
	logging.Info(fmt.Sprintf("%+v", found))
}
