package service

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PriceHistory struct {
	History []PriceItem `json:"history"`
}

type PriceMarket struct {
	MarketData []PriceItem `json:"marketdata"`
}

type PriceItem struct {
	TradeDate string  `json:"TRADEDATE"`
	Open      float32 `json:"OPEN"`
	Close     float32 `json:"CLOSE"`
	Low       float32 `json:"LOW"`
	High      float32 `json:"HIGH"`
	Last      float32 `json:"LAST"`
	Current   float32 `json:"CURRENTVALUE"`
}

func GetPriceHistoryByJson(data []byte) []PriceItem {
	prices := make([]PriceItem, 0, 100)
	h := []PriceHistory{{History: prices}}

	err := json.Unmarshal(data, &h)
	if err != nil {
		fmt.Println(err)
	}
	if len(h) > 1 {
		return h[1].History
	}
	return nil
}

func GetCurrentPriceByJson(data []byte) (*PriceItem, error) {
	prices := make([]PriceItem, 0, 100)
	h := []PriceMarket{{MarketData: prices}}

	err := json.Unmarshal(data, &h)
	if err != nil {
		return nil, err
	}
	if len(h) > 1 && len(h[1].MarketData) > 0 {
		return &h[1].MarketData[0], nil
	}
	return nil, errors.New("нет данных о цене")
}
