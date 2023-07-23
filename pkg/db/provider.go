package db

import (
	"github.com/ultra-supara/containerdev23/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	db *gorm.DB
}

func NewDatabaseProvider(dsn string) (*DatabaseProvider, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migration(db)
    return &DatabaseProvider{db: db}, nil
}

func NewDatabaseCon(conf config.Config) (*gorm.DB, error) {
    dbp, err := NewDatabaseProvider(conf.DatabaseConfig.DSN)
    if err != nil {
        return nil, err
    }
    return dbp.db, nil
}
