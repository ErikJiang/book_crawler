package parse

import (
	"log"
	"strconv"

	"github.com/JiangInk/book_crawler/model"
	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	Page int
	Url  string
}

func ParsePages(url string) (pages []Page) {
	log.Println("enter ParsePages.")
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	pages = append(pages, Page{Page: 1, Url: ""})

	doc.Find("div.indent > div.paginator > a").Each(func(index int, item *goquery.Selection) {
		log.Printf("> item: %v", item)
		page, _ := strconv.Atoi(item.Text())
		url, _ := item.Attr("href")
		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})
	return pages
}

func ParseBooks(doc *goquery.Document) (books []model.DoubanBook) {
	return books
}
