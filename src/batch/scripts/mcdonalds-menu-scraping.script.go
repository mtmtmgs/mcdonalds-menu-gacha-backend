package scripts

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/uptrace/bun"
)

const (
	morning = "朝マック"
	regular = "レギュラー"
	night   = "夜マック"
)

const (
	burger   = "バーガー"
	set      = "セット"
	sideMenu = "サイドメニュー"
	drink    = "ドリンク"
	happySet = "ハッピーセット"
	sweet    = "スイーツ"
	mcCafe   = "マックカフェ"
)

func FetchAndRegisterMcdonaldsMenu(db *bun.DB) error {
	fmt.Println("----------FetchAndRegisterMcdonaldsMenu() START----------")
	ch := make(chan models.Menu, 100)
	var wg sync.WaitGroup

	wg.Add(2)
	go fetchBurgerMenu(ch, &wg)
	go fetchSetMenu(ch, &wg)
	wg.Wait()
	close(ch)

	var menuList []models.Menu
	for menu := range ch {
		menuList = append(menuList, menu)
	}

	ctx := context.Background()
	_, err := db.NewInsert().Model(&menuList).Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Println("----------FetchAndRegisterMcdonaldsMenu() END----------")
	return nil
}

func fetchBurgerMenu(ch chan models.Menu, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get("https://www.mcdonalds.co.jp/menu/burger/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	selector := fmt.Sprintf(".collection-products [data-daypart='%s'] .product-list-wrapper .product-list", regular)
	doc.Find(selector).Each(func(i int, elem *goquery.Selection) {
		name := regexp.MustCompile("®").ReplaceAllString(elem.Find("strong").Text(), "")
		priceStr := regexp.MustCompile("~").ReplaceAllString(elem.Find(".product-list-card-price .product-list-card-price-number").Text(), "")
		priceNum, _ := strconv.ParseInt(priceStr, 10, 64)

		menu := models.Menu{
			Name:     name,
			Price:    priceNum,
			Category: burger,
			Hourly:   regular,
		}
		ch <- menu
	})
}

func fetchSetMenu(ch chan models.Menu, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get("https://www.mcdonalds.co.jp/menu/set/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	selector := fmt.Sprintf(".collection-products [data-daypart='%s'] .product-list-wrapper .product-list", regular)
	doc.Find(selector).Each(func(i int, elem *goquery.Selection) {
		name := regexp.MustCompile("®").ReplaceAllString(elem.Find("strong").Text(), "")
		priceStr := regexp.MustCompile("~").ReplaceAllString(elem.Find(".product-list-card-price .product-list-card-price-number").Text(), "")
		priceNum, _ := strconv.ParseInt(priceStr, 10, 64)

		menu := models.Menu{
			Name:     name,
			Price:    priceNum,
			Category: set,
			Hourly:   regular,
		}
		ch <- menu
	})
}
