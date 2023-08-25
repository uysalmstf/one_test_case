package main

import (
	"fmt"
	"log"
	"one_test_case/DBConfig"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBConfig.DB, err = gorm.Open("mysql", DBConfig.DbURL(DBConfig.BuildDBConfig()))
	fmt.Println(DBConfig.DB)
	if err != nil {
		panic(err)
	}

	defer DBConfig.DB.Close()
}
