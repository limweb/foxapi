package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	dB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

func SetupDB() {
	db := dB
	db, err = gorm.Open("sqlite3", "./webapi.db")
	if err != nil {
		DBErr = err
		fmt.Println("db err: ", err)
	}
	// Change this to true if you want to see SQL queries
	// ไว้แสดง Sql ออกทาง Log เมื่อต้องการ Debug คำสั่ง Sql
	//----------------Add MigrateDB --------------------------------
	dB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return dB
}

// GetDBErr helps you to get a connection
func GetDBErr() error {
	return DBErr
}
