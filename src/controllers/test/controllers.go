package test

import (
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
