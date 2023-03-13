package commands

import (
	"errors"
	"go-help-bot/internal/app/common"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PingParams struct {
	Url    string
	Count  int
	Period int
}

func (p *PingParams) PingSite(callback func(isAvailable bool)) {
	ticker := time.Tick(time.Second * time.Duration(p.Period))
	isAvailable := false
	i := 0
	for i < p.Count {
		<-ticker
		resp, _ := http.Get(p.Url)
		if resp.StatusCode == 200 {
			isAvailable = true
			common.DebugMsg("site available: " + p.Url)
			break
		} else {
			common.DebugMsg("ping: " + p.Url)
		}
		i++
	}
	callback(isAvailable)
}

func GetPingParamsByText(text string) (*PingParams, error) {
	p := &PingParams{Count: 100, Period: 10}
	params := strings.Split(text, " ")
	paramCount := len(params)
	if paramCount < 2 {
		return nil, errors.New("url parameter is required")
	}
	for i := 0; i < paramCount; i++ {
		switch i {
		case 1:
			p.Url = params[i]
		case 2:
			p.Count, _ = strconv.Atoi(params[i])
		case 3:
			p.Period, _ = strconv.Atoi(params[i])
		}
	}
	return p, nil
}
