package model

import "time"

type ItemBodegaAurrea struct {
	Data       Data       `json:"data"`
	Extensions Extensions `json:"extensions"`
}
type Categories struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}
type Products struct {
	ID                 string       `json:"id"`
	Name               string       `json:"name"`
	PhotosUrls         []string     `json:"photosUrls"`
	Sku                string       `json:"sku"`
	Unit               string       `json:"unit"`
	Price              int          `json:"price"`
	SpecialPrice       int          `json:"specialPrice"`
	Promotion          interface{}  `json:"promotion"`
	Stock              int          `json:"stock"`
	NutritionalDetails string       `json:"nutritionalDetails"`
	ClickMultiplier    int          `json:"clickMultiplier"`
	SubQty             int          `json:"subQty"`
	SubUnit            string       `json:"subUnit"`
	MaxQty             int          `json:"maxQty"`
	MinQty             int          `json:"minQty"`
	SpecialMaxQty      int          `json:"specialMaxQty"`
	Ean                []string     `json:"ean"`
	Boost              int          `json:"boost"`
	ShowSubUnit        bool         `json:"showSubUnit"`
	IsActive           bool         `json:"isActive"`
	Slug               string       `json:"slug"`
	Categories         []Categories `json:"categories"`
	Typename           string       `json:"__typename"`
}
type Paginator struct {
	Pages    int    `json:"pages"`
	Page     int    `json:"page"`
	Typename string `json:"__typename"`
}
type GetProducts struct {
	RedirectTo interface{} `json:"redirectTo"`
	Products   []Products  `json:"products"`
	Paginator  Paginator   `json:"paginator"`
	Typename   string      `json:"__typename"`
}
type Data struct {
	GetProducts GetProducts `json:"getProducts"`
}
type Resolvers struct {
	Path        []string `json:"path"`
	ParentType  string   `json:"parentType"`
	FieldName   string   `json:"fieldName"`
	ReturnType  string   `json:"returnType"`
	StartOffset int      `json:"startOffset"`
	Duration    int      `json:"duration"`
}
type Execution struct {
	Resolvers []Resolvers `json:"resolvers"`
}
type Tracing struct {
	Version   int       `json:"version"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Duration  int       `json:"duration"`
	Execution Execution `json:"execution"`
}
type Extensions struct {
	Tracing Tracing `json:"tracing"`
}
