package initializes

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	var err error
	dsn := "sqlserver://sa:29062004@localhost:1433?database=NCKH"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connect to database failed")
	}
}
