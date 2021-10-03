package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	model "gihub.com/dna2/promos/model"
	req "gihub.com/dna2/promos/request"
	src "gihub.com/dna2/promos/source"
)

type bodegaARepository struct {
	source      *src.Source
	url         string
	requestList []req.RequestBAurrera
}

func newInstanceBodegaARepo(source *src.Source) (*bodegaARepository, error) {
	if source == nil {
		return nil, fmt.Errorf("newInstanceSorianaRepo: Empty source")
	}

	requestHogar := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41471",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    500,
				CurrentPage: 1,
			},
		},
	}

	requestECompu := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41428",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestVinos := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41555",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestDespensa := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41443",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestHigiene := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41491",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestBebes := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41543",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestFarmacia := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41510",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestCongelados := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41438",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestCarnes := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41464",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestLacteos := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41556",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestBebidas := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41536",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestMascotas := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41419",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestTortillerias := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41525",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestSalchicha := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41532",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestFrutas := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41506",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	requestRopa := req.RequestBAurrera{
		Query: "query ($pagination: paginationInput, $search: SearchInput, $storeId: ID!, $categoryId: ID, $onlyThisCategory: Boolean, $filter: ProductsFilterInput, $orderBy: productsSortInput) {\n  getProducts(pagination: $pagination, search: $search, storeId: $storeId, categoryId: $categoryId, onlyThisCategory: $onlyThisCategory, filter: $filter, orderBy: $orderBy) {\n    redirectTo\n    products {\n      id\n      name\n      photosUrls\n      sku\n      unit\n      price\n      specialPrice\n      promotion {\n        description\n        type\n        isActive\n        conditions\n        __typename\n      }\n      stock\n      nutritionalDetails\n      clickMultiplier\n      subQty\n      subUnit\n      maxQty\n      minQty\n      specialMaxQty\n      ean\n      boost\n      showSubUnit\n      isActive\n      slug\n      categories {\n        id\n        name\n        __typename\n      }\n      __typename\n    }\n    paginator {\n      pages\n      page\n      __typename\n    }\n    __typename\n  }\n}\n",
		Variables: req.Variables{
			CategoryID:       "41411",
			StoreID:          "565",
			OnlyThisCategory: false,
			Pagination: req.Pagination{
				PageSize:    1000,
				CurrentPage: 1,
			},
		},
	}

	reqList := []req.RequestBAurrera{
		requestHogar,
		requestECompu,
		requestVinos,
		requestDespensa,
		requestHigiene,
		requestBebes,
		requestFarmacia,
		requestCongelados,
		requestCarnes,
		requestLacteos,
		requestBebidas,
		requestMascotas,
		requestTortillerias,
		requestSalchicha,
		requestFrutas,
		requestRopa,
	}
	return &bodegaARepository{source: source, url: "https://deadpool.instaleap.io/api/v2", requestList: reqList}, nil
}

func (repo *bodegaARepository) StartScraping(fOnFinish FnOnFinish) {
	for _, request := range repo.requestList {
		log.Printf("BodegaAurrea - StartScraping: %s - %s\n", repo.url, request.Variables.CategoryID)
		itemResponse, err := repo.callService(&request)
		if err != nil {
			continue
		}
		repo.iteratePages(&request, itemResponse)
	}
	fOnFinish()
}

func (repo *bodegaARepository) callService(request *req.RequestBAurrera) (*model.ItemBodegaAurrea, error) {
	jsonReq, _ := json.Marshal(request)
	resp, err := http.Post(repo.url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("Error call: %s,%v", repo.url, err.Error())
	}
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//responseString := string(responseBytes)
	var responseItem model.ItemBodegaAurrea
	json.Unmarshal(responseBytes, &responseItem)

	return &responseItem, nil
}

func (repo *bodegaARepository) iteratePages(request *req.RequestBAurrera, responseItem *model.ItemBodegaAurrea) {
	for _, item := range responseItem.Data.GetProducts.Products {
		specialPrice := item.Price
		diff := 0
		deal := 0
		if item.SpecialPrice > 0 {
			specialPrice = item.SpecialPrice
			diff = item.Price - specialPrice
			deal++
		}

		if item.Promotion != nil {
			deal++
		}

		itemPromo := model.ItemPromo{
			Chain:     1,
			Name:      item.Name,
			Link:      item.Slug,
			BasePrice: float64(item.Price),
			Price:     float64(specialPrice),
			Diff:      float64(diff),
			Deal:      uint8(deal),
			Date:      time.Now(),
			Sku:       item.Sku,
		}

		if len(item.Categories) > 0 {
			itemPromo.Category = item.Categories[0].Name
		}

		repo.source.Broker.SendMsg(itemPromo)
	}

	if responseItem.Data.GetProducts.Paginator.Pages > responseItem.Data.GetProducts.Paginator.Page {
		request.Variables.Pagination.CurrentPage = responseItem.Data.GetProducts.Paginator.Page + 1
		log.Printf("BodegaAurrea - Category:%s - Page:%d\n", request.Variables.CategoryID, request.Variables.Pagination.CurrentPage)
		itemResponse, err := repo.callService(request)
		if err == nil {
			repo.iteratePages(request, itemResponse)
		}
	}
}
