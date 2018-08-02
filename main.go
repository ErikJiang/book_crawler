package main

import (
	"log"

	"github.com/JiangInk/book_crawler/model"
	"github.com/JiangInk/book_crawler/parse"
)

var BaseURL = "https://book.douban.com/top250"

func start() {
	pages := parse.ParsePages(BaseURL)
	log.Printf("pages: %v", pages)
}

func main() {

	start()

	defer model.DB.Close()

}
