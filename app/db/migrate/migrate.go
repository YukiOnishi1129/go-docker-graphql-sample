package main

import (
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/db"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/entities"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{}, &entities.Todo{})
}

func main() {
	dbCon := db.Init()
	// dbを閉じる
	defer db.CloseDB(dbCon)

	migrate(dbCon)
}