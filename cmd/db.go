package cmd

import (
	"errors"
	"math"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
