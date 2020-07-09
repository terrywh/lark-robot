// +build !debug

package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/terrywh/lark-robot/v1/client"
)

func init() {
	app_id, _ := os.LookupEnv("APP_ID")
	app_secret, _ := os.LookupEnv("APP_SECRET")
	c = client.New(app_id, app_secret)
	e = echo.New()
	address, _ = os.LookupEnv("BIND_ADDR")
}
