package main

import (
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/config"
	"github.com/mirzha99/FP-BDS-Sanbercode-Go-52-Sistem-Pendataan-Penduduk/router"
)

func main() {
	//call .env
	config.Loadenv()
	//call database gorm driver mysql
	config.ConnectionDB()
	//run gin framework
	router.Router()
}
