package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type LarkClient struct {
	httpCli      *http.Client
	app_id       string
	app_secret   string
	access_token string
}

func New(app_id, app_secret string) *LarkClient {
	cli := &LarkClient{
		&http.Client{Timeout: 3 * time.Second}, app_id, app_secret, "",
	}
	go func() {
		for {
			cli.RefreshAccessToken()
			time.Sleep(time.Second * 1800)
		}
	}()

	return cli
}

func (c *LarkClient) prepareRequest(method, url string, body interface{}) (*http.Request, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.access_token)
	return req, err
}

func (c *LarkClient) sendForResponse(req *http.Request, i interface{}) error {
	res, err := c.httpCli.Do(req)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()
	return json.Unmarshal(data, i)
}

type tenant_access_token_request_t struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}
type tenant_access_token_response_t struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Token   string `json:"tenant_access_token"`
}

func (c *LarkClient) RefreshAccessToken() error {
	var rq tenant_access_token_request_t
	rq.AppId = c.app_id
	rq.AppSecret = c.app_secret
	data, err := json.Marshal(rq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/", bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	var rs tenant_access_token_response_t
	c.sendForResponse(req, &rs)
	c.access_token = "Bearer " + rs.Token
	return nil
}

type RichTextMessageContent struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type richTextMessageContent struct {
	Text string `json:"text"`
}

type richTextMessage struct {
	ChatId      string                 `json:"chat_id"`
	RootId      string                 `json:"root_id"`
	MessageType string                 `json:"msg_type"`
	Content     richTextMessageContent `json:"content"`
}

type default_response_t struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Data    map[string]interface{} `json:"data"`
}

func (c *LarkClient) MessageSend(chat_id, root_id, content string) (string, error) {
	if c == nil {
		fmt.Println("> MessageSend: ", chat_id, root_id, content)
		return "om_fake_message_id", nil
	}
	var msg richTextMessage
	msg.ChatId = chat_id
	msg.RootId = root_id
	msg.MessageType = "text"
	msg.Content.Text = content

	req, err := c.prepareRequest("POST", "https://open.feishu.cn/open-apis/message/v4/send/", msg)
	if err != nil {
		return "", err
	}

	var rs default_response_t
	c.sendForResponse(req, &rs)
	if rs.Code != 0 {
		return "", errors.New(rs.Message)
	}
	return rs.Data["message_id"].(string), nil
}
