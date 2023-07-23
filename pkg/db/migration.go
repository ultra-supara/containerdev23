package db

import (
	"github.com/ultra-supara/containerdev23/pkg/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&entity.Book{})
}
