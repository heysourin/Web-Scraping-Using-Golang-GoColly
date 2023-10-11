package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
)

type Item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"img-url"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	var items []Item

	c.OnHTML("article.product_pod", func(h *colly.HTMLElement) {
		item := Item{
			Name:   h.ChildText("h3"),
			Price:  h.ChildText("p.price_color"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)
	})

	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://books.toscrape.com/catalogue/category/books/mystery_3/index.html")

	// Print the populated items
	fmt.Println(items)

	content, err := json.Marshal(items)
	err = ioutil.WriteFile("./products.json", content, 0644)

	if err != nil {
		panic(err)
	}
}

/*//TODO
1. All the details(values) are loading inside one json key.Done --> You can't give the parent tag, that contains all the files. You have to choose such a parent-tag that is repeating itself.


2. 'Next page' is not happening. Done

//!
Idea 1: Will try without wg - timeout
*/
