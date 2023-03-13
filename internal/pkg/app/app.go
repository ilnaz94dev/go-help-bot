package app

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go-help-bot/internal/app/common"
	"go-help-bot/internal/app/endpoint"
	"net/http"
)

type App struct {
}

func Run() {
	initCore()

	err := endpoint.StartTlgBot(nil)
	if err != nil {
		common.HandleError("Telegram error", err)
	}
}

func initCore() {
	if err := godotenv.Load(); err != nil {
		common.HandleError("go-dotenv not loaded", err)
	}
}

func initRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/start_bot", endpoint.StartTlgBot)
	e.Match([]string{http.MethodGet, http.MethodPost}, "/check", endpoint.CheckSite)
}
