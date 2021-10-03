package repo

import (
	"fmt"
	"log"
	"strings"
	"time"

	model "gihub.com/dna2/promos/model"
	src "gihub.com/dna2/promos/source"
	"gihub.com/dna2/promos/util"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/gocolly/colly"
)

type elektraRepository struct {
	source       *src.Source
	urlContext   string
	initEndpoint string
}

func newInstanceElektraRepo(source *src.Source) (*elektraRepository, error) {
	if source == nil {
		return nil, fmt.Errorf("newInstanceElektraRepo: Empty source")
	}
	return &elektraRepository{source: source, urlContext: "https://www.elektra.com.mx", initEndpoint: "/linea-blanca/boilers?map=category-1,category-2,category-2,category-2,category-2,category-2&query=/linea-blanca/boilers/cocina/electrodomesticos/lavado-y-secado/refrigeradores-y-congeladores&searchState"}, nil
}

func (repo *elektraRepository) StartScraping(fOnFinish FnOnFinish) {
	log.Println("Elektra: StartScraping")
	fullPage := fmt.Sprintf("%s%s", repo.urlContext, repo.initEndpoint)

	geziyor.NewGeziyor(&geziyor.Options{
		LogDisabled: true,
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(fullPage, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			//fmt.Println(string(r.Body))
			r.HTMLDoc.Find(`section`).Each(func(_ int, s *goquery.Selection) {
				log.Println(s.Text())
			})
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
	fOnFinish()
}

func (repo *elektraRepository) findArticle(e *colly.HTMLElement) *model.ItemPromo {
	itemPromo := &model.ItemPromo{
		Chain: 0,
		Date:  time.Now(),
	}

	nodeContentTitle := e.DOM.Find("div.wrap-text-hook.wrap-text.plp-grid-word-break").First()
	if nodeContentTitle.Length() <= 0 {
		return nil
	}

	//TITLE AND LINK
	attrTitle, _ := nodeContentTitle.Children().First().Attr("title")
	attrHref, _ := nodeContentTitle.Children().First().Attr("href")
	hrefClean := strings.TrimSpace(attrHref)
	nameClean := strings.TrimSpace(attrTitle)
	if len(hrefClean) <= 0 && len(nameClean) <= 0 {
		return nil
	}
	itemPromo.Link = fmt.Sprintf("%s%s", repo.urlContext, hrefClean)
	itemPromo.Name = nameClean

	//Category
	listCategory := strings.Split(hrefClean, "/")
	if len(listCategory) > 3 {
		itemPromo.Category = listCategory[3]
	}

	//PRICE
	nodePrice := e.DOM.Find("div.product__list--price-panel").First()
	if nodePrice.Length() <= 0 {
		return nil
	}

	basePriceClean := strings.TrimSpace(nodePrice.Children().Eq(0).Text())
	priceClean := strings.TrimSpace(nodePrice.Children().Eq(1).Text())
	//log.Printf("PriceBase:%s - PriceFinal: %s",basePriceClean,priceClean)

	if len(basePriceClean) > 0 {
		basePriceCleanFormat, err := util.ParseMoney(&basePriceClean)
		if err == nil {
			itemPromo.BasePrice = basePriceCleanFormat
		}
	}

	if len(priceClean) > 0 {
		priceCleanFormat, err := util.ParseMoney(&priceClean)
		if err == nil {
			itemPromo.Price = priceCleanFormat
		}
	}

	if itemPromo.BasePrice == 0.0 {
		itemPromo.BasePrice = itemPromo.Price
	}

	if itemPromo.Price < itemPromo.BasePrice {
		itemPromo.Deal++
	}

	//DIFF
	itemPromo.Diff = itemPromo.BasePrice - itemPromo.Price

	//ISPROMO
	nodePromo := e.DOM.Find("img.PLP-promotion-icon").First()
	if nodePromo.Length() > 0 {
		itemPromo.Deal++
	}

	return itemPromo
}
