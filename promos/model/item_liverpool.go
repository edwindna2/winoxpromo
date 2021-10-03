package model

type ItemLiverpool struct {
	Links           []interface{}     `json:"links"`
	CarouselContent []CarouselContent `json:"carouselContent"`
	Status          Status            `json:"status"`
}
type InnerFilter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type BuryStrata struct {
	InnerFilter InnerFilter `json:"innerFilter"`
}
type Records struct {
	IsMarketPlace         string `json:"isMarketPlace"`
	ProductThumbnailImage string `json:"product.thumbnailImage"`
	SkuLargeImage         string `json:"sku.largeImage"`
	GroupType             string `json:"groupType"`
	MaximumListPrice      string `json:"maximumListPrice"`
	SkuSalePrice          string `json:"sku.sale_Price"`
	MinimumListPrice      string `json:"minimumListPrice"`
	ProductBrand          string `json:"product.brand"`
	ProductDisplayName    string `json:"product.displayName"`
	SellerName            string `json:"sellerName"`
	NumRecords            string `json:"numRecords"`
	ProductCategory       string `json:"product.category"`
	ProductLargeImage     string `json:"product.largeImage"`
	ProductRepositoryID   string `json:"product.repositoryId"`
	SkuPromoPrice         string `json:"sku.promoPrice"`
	SkuThumbnailImage     string `json:"sku.thumbnailImage"`
	IsHybridProduct       string `json:"isHybridProduct"`
	MinimumPromoPrice     string `json:"minimumPromoPrice"`
	SkuListPrice          string `json:"sku.list_Price"`
	SkuRepositoryID       string `json:"sku.repositoryId"`
	MaximumPromoPrice     string `json:"maximumPromoPrice"`
	Seller                string `json:"Seller"`
	ProductType           string `json:"productType"`
}
type CarouselContent struct {
	BuryStrata              []BuryStrata  `json:"buryStrata"`
	BoostStrata             []interface{} `json:"boostStrata"`
	DefaultCategory         string        `json:"defaultCategory"`
	Records                 []Records     `json:"records"`
	Type                    string        `json:"@type"`
	CarouselPath            string        `json:"carouselPath"`
	ShowMoreLink            bool          `json:"showMoreLink"`
	DimensionName           string        `json:"dimensionName"`
	MoreLinkStaticTextField string        `json:"moreLinkStaticTextField"`
	OverrideActiveCategory  bool          `json:"overrideActiveCategory"`
	DimensionID             string        `json:"dimensionId"`
	MoreLinkText            string        `json:"moreLinkText"`
	TotalNumRecs            int           `json:"totalNumRecs"`
	MaximumPrice            string        `json:"maximumPrice"`
	MaximumNumRecords       int           `json:"maximumNumRecords"`
	Name                    string        `json:"name"`
	MinimumPrice            string        `json:"minimumPrice"`
	MinNumRecords           int           `json:"minNumRecords"`
}
type Status struct {
	Status string `json:"status"`
}
