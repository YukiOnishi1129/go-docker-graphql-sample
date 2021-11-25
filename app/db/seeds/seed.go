package main

import (
	"fmt"
	"strconv"

	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/db"
	"github.com/YukiOnishi1129/go-docker-graphql-sample/app/db/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func userSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := entities.User {
			Name: "ユーザー"+strconv.Itoa(i+1),
			Email: "sample"+strconv.Itoa(i+1)+"@gmail.com",
			Password: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}


func todoSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		var userId int
		if i < 5 {
			userId = 1
		} else {
			userId = 2
		}
		todo := entities.Todo {
			Title: "タイトル"+strconv.Itoa(i+1),
			Comment: "コメント"+strconv.Itoa(i+1),
			UserId: userId,
		}

		if err := db.Create(&todo).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}


func main() {
	dbCon := db.Init()
	// dBを閉じる
	defer db.CloseDB(dbCon)

	if err := userSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}

	if err := todoSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
        return
	}
}