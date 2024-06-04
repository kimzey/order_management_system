package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

func NewPostgresDatabase() Database {
	dsn := "host=localhost port=5433 user=postgres password=123456 dbname=orderdb sslmode=disable search_path=public "

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	fmt.Println("Connect Database : ", conn.Name())

	return &postgresDatabase{conn}
}

func (p *postgresDatabase) Connect() *gorm.DB {
	return p.DB
}
