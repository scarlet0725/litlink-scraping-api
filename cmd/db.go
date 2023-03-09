package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/scarlet0725/prism-api/ent"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func InitDB(conf DBConfig) (*ent.Client, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	c := mysql.Config{
		DBName:    conf.Database,
		User:      conf.User,
		Passwd:    conf.Password,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		ParseTime: true,
		Loc:       jst,
	}

	for i := 0; ; i++ {
		if i > 10 {
			return nil, errors.New("failed to connect database")
		}

		db, err := sql.Open("mysql", c.FormatDSN())

		if err != nil {
			wait(i)
			continue
		}

		db.SetMaxOpenConns(10)

		err = db.Ping()

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

func getDBConfig() DBConfig {
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConf := DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	}

	return dbConf
}
