package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"example.com/GoMigrate/dao/model"
)

func main() {
  dsn := "cognusboi:password123@tcp(0.0.0.0:3306)/gomigratedb?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  var person model.PERSON
  db.First(&person, 1)
  fmt.Printf("%s\n", person.FirstName)
}
