package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/terrywh/lark-robot/v1/client"
	"github.com/terrywh/lark-robot/v1/toolkit"
)

type event_callback_t struct {
	Type      string                  `json:"type"`
	Token     string                  `json:"token,omitempty"`
	Challenge string                  `json:"challenge,omitempty"`
	Event     toolkit.TextMessageChat `json:"event,omitempty"`
}

type challenge_response_t struct {
	Challenge string `json:"challenge"`
}

var c *client.LarkClient
var e *echo.Echo
var address string

func main() {
	// 初始化过程条件编译
	e.POST("/", func(cc echo.Context) error {
		var req event_callback_t
		cc.Bind(&req)
		if req.Type == "event_callback" {

			if req.Event.Type == "message" && (req.Event.MessageType == "text" || req.Event.MessageType == "post") {
				toolkit.New(c).Do(&req.Event)
				return cc.String(200, "RECEIVED: "+req.Event.OpenMessageId)
			}
			return cc.String(404, "UNKNOWN")
		} else if req.Type == "url_verification" {
			var res challenge_response_t
			res.Challenge = req.Challenge
			return cc.JSON(200, res)
		} else {
			return cc.String(404, "UNKNOWN")
		}
	})
	e.Use(middleware.Logger())
	e.Start(address)
}
