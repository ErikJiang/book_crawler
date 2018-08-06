package parse

import (
	"log"
	"strconv"
	"strings"

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

	pages = append(pages, Page{Page: 1, Url: "https://book.douban.com/top250"})

	doc.Find("div.indent > div.paginator > a").Each(func(index int, item *goquery.Selection) {
		page, _ := strconv.Atoi(item.Text())
		url, _ := item.Attr("href")
		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})
	return pages
}

func ParseBooks(url string) (books []model.DoubanBook, err error) {
	log.Printf("enter ParseBooks url: %s", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("div.indent > table").Each(func(index int, bookItem *goquery.Selection) {
		head := bookItem.Find("tbody > tr.item > td")
		title, _ := head.Eq(1).Find("div.pl2 > a").Attr("title")
		subtitle := head.Eq(1).Find("div.pl2 > span").Text()
		infos := strings.Split(head.Eq(1).Find("p.pl").Text(), " / ")

		var author, translator, publisher, pubDate, price string
		if len(infos) == 4 {
			author = infos[0]
			publisher = infos[1]
			pubDate = infos[2]
			price = infos[3]

		} else if len(infos) == 5 {
			author = infos[0]
			translator = infos[1]
			publisher = infos[2]
			pubDate = infos[3]
			price = infos[4]
		} else {
			log.Fatalf("infos str: %s", infos)
		}

		star := head.Eq(1).Find("div.star > span.rating_nums").Text()
		comments := head.Eq(1).Find("div.star > span.pl").Text()
		comments = strings.Replace(comments, "\n", "", -1)
		comments = strings.Replace(comments, " ", "", -1)
		prefix := "("
		suffix := "人评价)"
		comments = strings.TrimPrefix(comments, prefix)
		comments = strings.TrimSuffix(comments, suffix)
		quote := head.Eq(1).Find("p.quote > span.inq").Text()

		book := model.DoubanBook{
			Title:      title,
			Subtitle:   subtitle,
			Author:     author,
			Translator: translator,
			Publisher:  publisher,
			PubDate:    pubDate,
			Price:      price,
			Star:       star,
			CommentNum: comments,
			Quote:      quote,
		}
		log.Printf("index: %d, item: %v", index, book)
		books = append(books, book)
	})

	return books, nil
}
