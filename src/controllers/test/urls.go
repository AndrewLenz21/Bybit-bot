package test

import (
	"github.com/labstack/echo/v4"
)

// Start router
func TestController(e *echo.Echo) {
	e.GET("/name", Yallo)
	e.GET("/getLastPnl", GetLastPnl)
	e.GET("/getOpenPositions", GetOpenPositions)
	e.GET("/getOpenOrders", GetOpenOrders)
}
