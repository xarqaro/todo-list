package database

import (
	"fmt"
	"log"
	"os"

	"github.com/MagaSu/todo-list/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	// Cоздаем URL для соединения с бд.
	// Имя пользователя базы данных, пароль и имя бд
	// берутся из переменных окружения
	// они описаны в файле .env

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Создаем подключение к бд
	// В &gorm.Config настраивается логгер,
	// который будет сохранять информацию
	// обо всех активностях с базой данных.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(1)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migration")
	db.AutoMigrate(&models.Task{})

	DB = Dbinstance{
		Db: db,
	}
}
