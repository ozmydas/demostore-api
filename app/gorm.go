package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/db_demo")

	if err != nil {
		log.Panicf("Failed Connect to DB : %v\n", err)
	}

	db.DB()
	// db.AutoMigrate($models.Notes{})
	return db
}
