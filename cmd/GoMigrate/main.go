package main

import (
	"fmt"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example.com/GoMigrate/dao/model"
)

func migrate(from *gorm.DB, to *gorm.DB, t reflect.Type, bufsize int) {
	var count int64
	elem := reflect.New(t).Interface()
	from.Model(elem).Count(&count)

	fmt.Printf("Count: %d\n", count)

	to.AutoMigrate(t)

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
	// mySQL connection
	dsn := "cognusboi:password123@tcp(0.0.0.0:3306)/gomigratedb?charset=utf8mb4&parseTime=True&loc=Local"
	mysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Postgres connection
	dsn = "host=0.0.0.0 user=cognusboi password=password123 dbname=gomigratedb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	postgres.AutoMigrate(&model.PERSON{})

	migrate(mysql, postgres, reflect.TypeOf(model.PERSON{}), 10)
}
