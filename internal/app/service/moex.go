package service

import (
	"time"
)

const moexApiUrl = "http://iss.moex.com/iss" // /history

type MoexRequestParams struct {
	isHistory bool
	engine    string
	market    string
	board     string
	security  string
	dateFrom  string
	dateTo    string
}

func GetTickerParams(secId string) *MoexRequestParams {
	return &MoexRequestParams{
		engine:   "stock",
		market:   "shares",
		board:    "TQBR",
		security: secId,
	}
}

func GetCurrencyParams(secId string) *MoexRequestParams {
	return &MoexRequestParams{
		engine:   "currency",
		market:   "index",
		board:    "FIXI",
		security: secId,
	}
}

func (p *MoexRequestParams) BuildUrl() string {
	url := moexApiUrl
	if p.isHistory {
		url += "/history"
	}
	if p.engine != "" {
		url += "/engines/" + p.engine
	}
	if p.engine != "" {
		url += "/markets/" + p.market
	}
	if p.engine != "" {
		url += "/boards/" + p.board
	}
	if p.engine != "" {
		url += "/securities/" + p.security
	}
	url += ".json?iss.meta=off&iss.json=extended"
	if p.dateFrom != "" {
		url += "&from=" + p.dateFrom
	}
	if p.dateTo != "" {
		url += "&till=" + p.dateTo
	}
	return url
}

func (p *MoexRequestParams) SetDateFrom(t time.Time) *MoexRequestParams {
	p.dateFrom = t.Format(time.DateOnly)
	return p
}

func (p *MoexRequestParams) SetDateTo(t time.Time) *MoexRequestParams {
	p.dateTo = t.Format(time.DateOnly)
	return p
}
