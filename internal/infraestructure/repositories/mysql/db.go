package mysql

import (
	"log/slog"
	"os"

	"github.com/maoudev/todo/internal/pkg/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB

	tables = []interface{}{
		&domain.User{},
		&domain.Task{},
	}
)

func connect() *gorm.DB {
	var err error
	dsn := os.Getenv("DSN")

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
	}

	migrate()

	slog.Info("Connected to the database!")
	return db
}

func migrate() {
	for _, t := range tables {
		if err := db.AutoMigrate(t); err != nil {
			slog.Error(err.Error())
			return
		}
	}
}
