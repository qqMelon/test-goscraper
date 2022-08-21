package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"goscraper/models"
	"goscraper/utils"
	"os"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	var items []models.Item

	c.OnHTML("div[itemprop=itemListElement]", func(el *colly.HTMLElement) {
		item := models.Item{
			Name:   el.ChildText("h2.product-title"),
			Price:  el.ChildText("div.sale-price"),
			ImgUrl: el.ChildAttr("img", "src"),
		}

		items = append(items, item)
	})

	c.OnHTML("[title=Next]", func(el *colly.HTMLElement) {
		nextPage := el.Request.AbsoluteURL(el.Attr("href"))
		err := c.Visit(nextPage)
		utils.ErroHandler(err)
	})

	c.OnRequest(func(req *colly.Request) {
		fmt.Println(req.URL.String())
	})

	err := c.Visit("https://j2store.net/demo/index.php/shop")
	utils.ErroHandler(err)

	content, err := json.Marshal(items)
	utils.ErroHandler(err)

	err = os.WriteFile("output/product.json", content, 0644)
	utils.ErroHandler(err)
}
