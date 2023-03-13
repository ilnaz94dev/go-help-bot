package commands

import (
	"go-help-bot/internal/app/service"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GetCurrency(code string) (string, error) {
	url := service.GetCurrencyParams(code + "FIX").SetDateFrom(time.Now()).BuildUrl()
	return getPriceByUrl(url)
}

func GetTicker(code string) (string, error) {
	url := service.GetTickerParams(code).SetDateFrom(time.Now()).BuildUrl()
	return getPriceByUrl(url)
}

func getPriceByUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	item, err := service.GetCurrentPriceByJson(data)
	if err != nil {
		return "", err
	}
	res := item.Current
	if res == 0 {
		res = item.Last
	}
	// TODO тестыы
	return strconv.FormatFloat(float64(res), 'f', 1, 32), nil
}
