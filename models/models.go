package models

import (
	"az-edu/library/db"
	"github.com/op/go-logging"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var logger = logging.MustGetLogger("models")

func CreateTable() {
	db.DB.DropTableIfExists(&Quesiton{})
	create := db.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	create.CreateTable(&Quesiton{})
}

func MigrateTable() {
	create := db.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	create.AutoMigrate(&Quesiton{})
}
