package main

import (
	"log"

	"github.com/JiangInk/book_crawler/model"
	"github.com/JiangInk/book_crawler/parse"
)

var BaseURL = "https://book.douban.com/top250"

func booksDataCrawler() {
	var doubanBooks []model.DoubanBook

	pages := parse.ParsePages(BaseURL)
	log.Printf("pages: %v", pages)
	for _, page := range pages {
		books, err := parse.ParseBooks(page.Url)
		if err != nil {
			log.Fatal(err)
		}
		doubanBooks = append(doubanBooks, books...)
	}
	log.Printf("doubanBooks: %v", doubanBooks)
	for index, book := range doubanBooks {
		if err := model.DB.Create(&book).Error; err != nil {
			log.Fatalf("index: %d, err: %v", index, err)
		}
	}
}

func main() {

	booksDataCrawler()

	defer model.DB.Close()

}
