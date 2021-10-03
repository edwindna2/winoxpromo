package util

import (
	"fmt"
	"strconv"
	"strings"
)

//ParseMoney Function
func ParseMoney(money *string) (float64, error) {

	if money == nil {
		return -1, fmt.Errorf("Invalid string: ","emtpy money")
	}

	priceClean := strings.ReplaceAll(*money, "$", "")
	priceClean = strings.ReplaceAll(priceClean, ",", "")
	floatValue, err := strconv.ParseFloat(priceClean, 64)
	if err != nil {
		return -1, err
	}

	return floatValue, nil
}
