package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"simpleblog/app/support"
)

func OpenDB() gorm.DB {
	db, err := gorm.Open("postgres", "dbname=gobone_dev sslmode=disable")
	if err != nil {
		support.LogStacktrace(err)
		panic(err)
	}

	//defaults
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()
	return db
}
