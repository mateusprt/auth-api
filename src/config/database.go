package config

import (
	"log"
	"os"

	"github.com/mateusprt/auth-api/src/models"
	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	gotenv.Load()
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	// "usuario:senha@tcp(127.0.0.1:porta)/nome_database?charset=utf8&parseTime=True&loc=Local"
	stringConnection := db_user + ":" + db_password + "@tcp(127.0.0.1:3306)/" + db_name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(stringConnection), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect database")
		log.Panic(err)
	}

	db.AutoMigrate(&models.User{})

	log.Println("Database connected successfully")

	return db
}
