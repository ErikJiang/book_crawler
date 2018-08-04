package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB       *gorm.DB
	username string = "test"
	password string = "jiangink"
	dbName   string = "book_crawler"
)

type DoubanBook struct {
	gorm.Model
	Title      string
	Subtitle   string
	Author     string
	Translator string
	Publisher  string
	PubDate    string
	Price      string
	Star       string
	CommentNum string
	Quote      string
}

func init() {
	var err error
	DB, err = gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
			username,
			password,
			dbName,
		),
	)
	if err != nil {
		log.Fatalf("gorm.Open.err: %v", err)
	}
	DB.SingularTable(true)
	DB.AutoMigrate(&DoubanBook{})
}
