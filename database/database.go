package database

import (
	"tsuruev/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "host=localhost user=postgres password=4444 dbname=backand-for-intucode port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	DB = db
	db.AutoMigrate(&models.Student{}, &models.Group{}, &models.Note{})
}
