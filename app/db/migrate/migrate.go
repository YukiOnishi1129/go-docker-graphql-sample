package main

import (
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/db"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Todo{})
}

func main() {
	dbCon := db.Init()
	// dbを閉じる
	defer db.CloseDB(dbCon)

	migrate(dbCon)
}