package test

import (
	bot_service "bybitbot/src/services/bot"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Yallo(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")

	if name == "" || age == "" {
		return c.String(http.StatusBadRequest, "error")
	}

	message := fmt.Sprintf("Tu nombre es %s y tienes %s años", name, age)
	fmt.Println(message)
	return c.String(http.StatusOK, message)
}

func GetLastPnl(c echo.Context) error {
	symbol := c.QueryParam("symbol")
	if symbol == "" {
		return c.String(http.StatusBadRequest, "error")
	}
	last_pnl := bot_service.BybitGetLastPnl(symbol, context.Background())
	return c.String(http.StatusOK, last_pnl)
}

func GetOpenPositions(c echo.Context) error {
	symbol := c.QueryParam("symbol")
	side := c.QueryParam("side")
	if symbol == "" || side == "" {
		return c.String(http.StatusBadRequest, "error")
	}
	positions := bot_service.BybitGetPositions(symbol, context.Background())
	return c.JSON(http.StatusOK, positions)
}

func GetOpenOrders(c echo.Context) error {
	symbol := c.QueryParam("symbol")
	side := c.QueryParam("side")
	if symbol == "" || side == "" {
		return c.String(http.StatusBadRequest, "error")
	}
	BybitOpenOrders := bot_service.BybitGetOpenOrders(symbol, side, context.Background())
	return c.JSON(http.StatusOK, BybitOpenOrders)
}
