package cmd

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	c := 0
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		time.Sleep(time.Second)
		c += 1
		if c > 30 {
			log.Fatal(err)
		}

	}
	return db, nil
}
