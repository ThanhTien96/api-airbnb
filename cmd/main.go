package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ThanhTien96/airbnb/internal/api"
	"github.com/ThanhTien96/airbnb/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// intialize and connect db
	dsn := "host=0.0.0.0 user=postgres password=admin dbname=airbnb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// check flag and handling
	isCleanData := flag.Bool("c", false, "clean data")
	isGenerateData := flag.Bool("g", false, "generate data")
	isMigrate := flag.Bool("m", false, "migrate model from code to database")
	flag.Parse()

	if *isCleanData {
		utils.CleanData(db)
	}
	if *isMigrate {
		fmt.Println("ldslkjfsd")
		err := utils.MigrateData(db)
		if err != nil {
			log.Fatal("Error while migrating data from model to database ", err)
		}
		// os.Exit(0)
		log.Println("Successfully migrated data")
	}

	if *isGenerateData {
		// utils.CreateUserRole(db)
		utils.CreateUsers(db, utils.USER_LIMIT)
	}

	e := echo.New()

	apiV1 := e.Group("/v1/api")

	apiV1.GET("/userRole", api.ApiGetUserRole(db))

	apiV1.GET("/user", api.ApiGetUsers(db))

	e.Logger.Fatal(e.Start(":8080"))
}
