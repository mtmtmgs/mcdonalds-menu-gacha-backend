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
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/consts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"
	"github.com/uptrace/bun"
)

const (
	burgerUrl   = "https://www.mcdonalds.co.jp/menu/burger/"
	setUrl      = "https://www.mcdonalds.co.jp/menu/set/"
	sideUrl     = "https://www.mcdonalds.co.jp/menu/side/"
	drinkUrl    = "https://www.mcdonalds.co.jp/menu/drink/"
	happySetUrl = "https://www.mcdonalds.co.jp/menu/happyset/"
	dessertUrl  = "https://www.mcdonalds.co.jp/menu/dessert/"
	mcCafeUrl   = "https://www.mcdonalds.co.jp/menu/mccafe/"
)

type category struct {
	name string
	url  string
}

/*
メイン処理
各メニューを取得してDBに保存
*/
func FetchAndRegisterMcdonaldsMenu(db *bun.DB) error {
	log.Println("------------------------------------------------------------")
	log.Println("START FetchAndRegisterMcdonaldsMenu()")
	log.Println("------------------------------------------------------------")
	// 時間帯別リスト
	var mealTimeTypes = []string{
		consts.Morning,
		consts.Noon,
		consts.Night,
		consts.Regular,
	}
	// メニューカテゴリリスト
	var categories = []category{
		{name: consts.Burger, url: burgerUrl},
		{name: consts.Set, url: setUrl},
		{name: consts.Side, url: sideUrl},
		{name: consts.Drink, url: drinkUrl},
		{name: consts.HappySet, url: happySetUrl},
		{name: consts.Dessert, url: dessertUrl},
		{name: consts.McCafe, url: mcCafeUrl},
	}

	ch := make(chan models.Menu, 300)
	var wg sync.WaitGroup

	wg.Add(len(categories))
	for _, category := range categories {
		go fetchMenus(ch, &wg, mealTimeTypes, category)
	}
	wg.Wait()
	close(ch)

	var menuList []models.Menu
	for menu := range ch {
		menu, err := models.NewMenu(menu)
		if err != nil {
			log.Printf("追加できませんでした: %v %s", menu, err)
			continue
		}
		menuList = append(menuList, menu)
	}

	ctx := context.Background()
	_, err := db.NewInsert().Model(&menuList).Exec(ctx)
	if err != nil {
		return err
	}
	log.Printf("menusテーブルへの保存が完了しました. %d件", len(menuList))
	log.Println("------------------------------------------------------------")
	log.Println("END FetchAndRegisterMcdonaldsMenu()")
	log.Println("------------------------------------------------------------")
	return nil
}

/*
メニュー取得
*/
func fetchMenus(ch chan models.Menu, wg *sync.WaitGroup, mealTimeTypes []string, category category) {
	defer wg.Done()

	res, err := http.Get(category.url)
	if err != nil {
		log.Println("------------------------------------------------------------")
		log.Println("EXECUTE FAILURE!")
		log.Println("------------------------------------------------------------")
		log.Fatalf("DETAIL: %s %v", category.url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("------------------------------------------------------------")
		log.Println("EXECUTE FAILURE!")
		log.Println("------------------------------------------------------------")
		log.Fatalf("DETAIL: %s status code: %d %s", category.url, res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("------------------------------------------------------------")
		log.Println("EXECUTE FAILURE!")
		log.Println("------------------------------------------------------------")
		log.Fatalf("DETAIL: %s %v", category.url, err)
	}

	var count int
	for _, mealTimeType := range mealTimeTypes {
		// 時間帯別で該当するメニューは存在しないので飛ばす
		if mealTimeType == consts.Morning && (category.name == consts.Drink || category.name == consts.Dessert || category.name == consts.McCafe) {
			continue
		}
		if mealTimeType == consts.Noon && category.name != consts.Set {
			continue
		}
		if mealTimeType == consts.Night && (category.name == consts.Drink || category.name == consts.HappySet || category.name == consts.Dessert || category.name == consts.McCafe) {
			continue
		}

		// メニューに基づいて対象とするセレクタを分ける
		var selector string
		if category.name == consts.Burger || category.name == consts.Set || category.name == consts.Side || category.name == consts.HappySet {
			selector = fmt.Sprintf(".collection-products [data-daypart='%s'] .product-list-wrapper .product-list", mealTimeType)
			if mealTimeType == consts.Noon {
				selector = fmt.Sprintf(".collection-products [data-daypart='%s'] .product-list-wrapper .product-list", "ひるまック（平日限定）")
			}
		}
		if category.name == consts.Drink || category.name == consts.Dessert {
			selector = ".product-list-wrapper .product-list"
		}
		if category.name == consts.McCafe {
			selector = ".product-list-wrapper.flex.flex-wrap.pb-8 .product-list"
		}

		doc.Find(selector).Each(func(i int, elem *goquery.Selection) {
			name := regexp.MustCompile("®").ReplaceAllString(elem.Find("strong").Text(), "")
			priceStr := regexp.MustCompile("~").ReplaceAllString(elem.Find(".product-list-card-price .product-list-card-price-number").Text(), "")
			priceNum, _ := strconv.ParseInt(priceStr, 10, 64)

			menu := models.Menu{
				Name:         name,
				Price:        priceNum,
				Category:     category.name,
				MealTimeType: mealTimeType,
			}
			count++
			ch <- menu
		})
	}
	log.Printf("%sの抽出が完了しました. %d件", category.name, count)
}
