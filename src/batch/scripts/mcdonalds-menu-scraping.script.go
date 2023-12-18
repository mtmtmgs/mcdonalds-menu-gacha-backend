package scripts

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/uptrace/bun"
)

func FetchAndRegisterMcdonaldsMenu(db *bun.DB) error {
	res, err := http.Get("https://www.mcdonalds.co.jp/menu/burger/")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find(".collection-products [data-daypart='レギュラー'] .product-list-wrapper .product-list").Each(func(i int, elem *goquery.Selection) {
		name := elem.Find("strong").Text()
		price := elem.Find(".product-list-card-price .product-list-card-price-number").Text()
		fmt.Printf("Review %d: %s %s\n", i, name, price)
	})

	return nil
}
