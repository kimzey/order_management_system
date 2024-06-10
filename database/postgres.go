package database

import (
	"fmt"
	"github.com/kizmey/order_management_system/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

func NewPostgresDatabase(conf *config.Database) Database {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s ",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.DBName,
		conf.SSLMode,
		conf.Schema,
	)
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
