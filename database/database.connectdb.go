package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func Connectmydatabase() {
	connStr := "postgres://postgres:root@localhost/discovery?sslmode=disable"

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Now we are connected to POSTGRESQL DATABASE.")
}

func GetUserDB() *gorm.DB {
	return db
}

func Closemydatabase() {
	sqlcon, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlcon.Close()
	fmt.Println("db closed successfully")
}
