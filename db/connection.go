//
// connection.go
// Copyright (C) 2022 tinix <tinix@archlinux>
//
// Distributed under terms of the MIT license.
//

package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=tinix password=password dbname=gorm_development port=5432"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected.")
	}
}
