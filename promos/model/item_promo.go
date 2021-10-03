package model

import (
	"encoding/json"
	"time"
)

type ItemPromo struct {
	Chain     uint8     `json:"chain"`
	Name      string    `json:"name"`
	Link      string    `json:"link"`
	BasePrice float64   `json:"bprice"`
	Price     float64   `json:"price"`
	Diff      float64   `json:"diff"`
	Category  string    `json:"category"`
	Deal      uint8     `json:"deal"`
	Date      time.Time `json:"date"`
	Sku       string    `json:"sku"`
}

func (obj ItemPromo) ToJSON() (string, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (obj ItemPromo) ToBytes() ([]byte, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return result, nil
}
