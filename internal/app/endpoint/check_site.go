package endpoint

import (
	"github.com/labstack/echo/v4"
	"go-help-bot/internal/app/commands"
	"net/http"
)

func CheckSite(c echo.Context) error {
	url := c.QueryParam("url")
	p := commands.PingParams{Url: url, Count: 100, Period: 10}
	go p.PingSite(nil)

	return c.String(http.StatusOK, "Site check started.")
}
