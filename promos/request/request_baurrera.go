package request

type RequestBAurrera struct {
	Variables Variables `json:"variables"`
	Query     string    `json:"query"`
}

type Pagination struct {
	PageSize    int `json:"pageSize"`
	CurrentPage int `json:"currentPage"`
}

type Variables struct {
	CategoryID       string     `json:"categoryId"`
	OnlyThisCategory bool       `json:"onlyThisCategory"`
	Pagination       Pagination `json:"pagination"`
	StoreID          string     `json:"storeId"`
}
