package main

import (
	"gin-flemarket/infra"
	"gin-flemarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate database")
	}
}