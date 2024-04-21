package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"example.com/GoMigrate/dao/model"
)

type DBMS struct {
	System   string `json:"system"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func (d *DBMS) dsn() string {
	switch d.System {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Database)
	case "postgres":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", d.Host, d.Username, d.Password, d.Database, d.Port)
	case "sqlite":
		return fmt.Sprintf("%s.db", d.Database)
	default:
		panic("Invalid DBMS")
	}
}

func getDsnFromJson(filename string) (DBMS, DBMS) {
	file, err := os.Open(filename)
	if err != nil {
		panic("Failed to open file")
	}
	defer file.Close()

	var data struct {
		From DBMS `json:"from"`
		To   DBMS `json:"to"`
	}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		panic("Failed to decode json")
	}
	return data.From, data.To
}

func (p *DBMS) getConnection() *gorm.DB {
	if p.System == "mysql" {
		conn, err := gorm.Open(mysql.Open(p.dsn()), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return conn
	} else if p.System == "postgres" {
		conn, err := gorm.Open(postgres.Open(p.dsn()), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return conn
	} else if p.System == "sqlite" {
		conn, err := gorm.Open(sqlite.Open(p.dsn()), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return conn
	} else {
		panic("Invalid DBMS")
	}
}

func migrate(from *gorm.DB, to *gorm.DB, t reflect.Type, bufsize int) {
	var count int64
	elem := reflect.New(t).Interface()
	from.Model(elem).Count(&count)

	fmt.Printf("Count: %d\n", count)

	to.AutoMigrate(elem)

	for i := 0; i < int(count); i += bufsize {
		newArr := reflect.New(reflect.SliceOf(t)).Interface()
		from.Limit(bufsize).Offset(i).Find(newArr)

		result := to.Create(newArr)
		if result.Error != nil {
			panic("Failed to migrate")
		}
	}
}

func main() {
	FromDBMS, ToDBMS := getDsnFromJson("res/conf.json")

	// From connection
	from := FromDBMS.getConnection()

	// To connection
	to := ToDBMS.getConnection()

	for _, t := range model.GetStructs() {
		migrate(from, to, t, 10)
	}
}
