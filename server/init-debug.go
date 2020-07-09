// +build debug

package main

import (
	"github.com/labstack/echo/v4"
)

func init() {
	e = echo.New()
	address = ":8080"
}
