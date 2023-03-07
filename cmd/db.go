package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/scarlet0725/prism-api/ent"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func ConnectDB(dsn string) (*gorm.DB, error) {
	for i := 0; ; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		if i > 10 {
			return nil, errors.New("failed to connect database")
		}
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
	}
}

func InitDB(conf DBConfig) (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)

	for i := 0; ; i++ {
		if i > 10 {
			return nil, errors.New("failed to connect database")
		}

		db, err := sql.Open("mysql", dsn)

		if err != nil {
			wait(i)
			continue
		}

		db.SetMaxOpenConns(10)

		db.Ping()

		if err != nil {
			wait(i)
			continue
		}

		drv := entsql.OpenDB("mysql", db)
		return ent.NewClient(ent.Driver(drv)), nil

	}
}

func wait(i int) {
	time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
}
