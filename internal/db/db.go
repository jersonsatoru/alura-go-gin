package db

import (
	"github.com/jersonsatoru/alura-go-gin/internal/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=satoru password=satoru dbname=satoru port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("was not possible connect with db")
	}
	DB.AutoMigrate(&orm.Student{})
}
