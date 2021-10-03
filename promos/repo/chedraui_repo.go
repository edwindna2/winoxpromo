package repo

import (
	"fmt"
	"log"
	"strings"
	"time"

	model "gihub.com/dna2/promos/model"
	src "gihub.com/dna2/promos/source"
	"gihub.com/dna2/promos/util"
	"github.com/gocolly/colly"
)

type chedrauiRepository struct {
	source	*src.Source
	urlContext string
	initEndpoint string
}

func newInstanceChedrauiRepo(source *src.Source) (*chedrauiRepository, error) {
	if source == nil {
		return nil, fmt.Errorf("newInstanceChedrauiRepo: Empty source")
	}
	return &chedrauiRepository{source: source,urlContext:"https://www.chedraui.com.mx",initEndpoint:"/Departamentos/c/MC?sort=relevance&pageSize=48&q=%3Arelevance&toggleView=grid#"}, nil
}

func (repo *chedrauiRepository) StartScraping(fOnFinish FnOnFinish) {
	log.Println("Chedraui: StartScraping")
	isRunning := true

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML(`div[id=plp_display]`, func(e *colly.HTMLElement) {
		//log.Printf("%s - %s:%s", repo.urlContext, "OnHTML", "div[id=plp_display]")
		e.ForEach("li", func(_ int, li *colly.HTMLElement) {
			item := repo.findArticle(li)
			if item != nil {
				//log.Printf("Item: %v", item)
				repo.source.Broker.SendMsg(item)
			}
		})
	})

	// On every a element which has href attribute call callback
	c.OnHTML(`a[rel=next]`, func(e *colly.HTMLElement) {
		//log.Printf("%s - %s:%s", repo.urlContext, "OnHTML", "a[rel=next]")
		link := e.Attr("href")
		absoluteURL := e.Request.AbsoluteURL(link)
		c.Visit(absoluteURL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("%s - %s:%s", r.Request.URL, "OnError", err.Error())
	})

	c.OnScraped(func(r *colly.Response) {
		if isRunning {
			isRunning = false
			log.Printf("%s - %s", r.Request.URL, "OnScraped")
			fOnFinish()
		}
	})

	// Start scraping
	c.Visit(fmt.Sprintf("%s%s",repo.urlContext,repo.initEndpoint))
}

func (repo *chedrauiRepository) findArticle(e *colly.HTMLElement) *model.ItemPromo {
	itemPromo := &model.ItemPromo{
		Chain: 0,
		Date: time.Now(),
	}

	nodeContentTitle := e.DOM.Find("div.wrap-text-hook.wrap-text.plp-grid-word-break").First()
	if nodeContentTitle.Length() <= 0 {
		return nil
	}

	//TITLE AND LINK
	attrTitle, _ := nodeContentTitle.Children().First().Attr("title")
	attrHref, _ := nodeContentTitle.Children().First().Attr("href")
	hrefClean:= strings.TrimSpace(attrHref)
	nameClean := strings.TrimSpace(attrTitle)
	if len(hrefClean) <= 0 && len(nameClean) <= 0 {
		return nil
	}
	itemPromo.Link = fmt.Sprintf("%s%s",repo.urlContext,hrefClean)
	itemPromo.Name = nameClean

	//Category
	listCategory := strings.Split(hrefClean,"/")
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
		itemPromo.BasePrice  = itemPromo.Price
	}

	if itemPromo.Price < itemPromo.BasePrice{
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
