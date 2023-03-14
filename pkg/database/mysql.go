package database

import (
	"fmt"
	"log"
	"wide_technologies/config"
	student "wide_technologies/internal/student/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(conf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func MigrateMysql(db *gorm.DB) {
	db.AutoMigrate(&student.Student{})
}
