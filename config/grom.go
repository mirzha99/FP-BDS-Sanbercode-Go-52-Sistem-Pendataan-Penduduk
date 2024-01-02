package config

import (
	"fmt"
	"os"

	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/models/Muser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	defer func() {
		recover := recover()
		if recover != nil {
			fmt.Println(recover)
			os.Exit(0)
		}
	}()
	dns := os.Getenv("dns_db")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Muser.User{})
	DB = db
}
