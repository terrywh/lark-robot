package toolkit

import (
	"regexp"
	"strings"

	"github.com/terrywh/lark-robot/v1/client"
)

type TextMessageChat struct {
	Type          string `json:"type"`
	MessageType   string `json:"msg_type"`
	OpenMessageId string `json:"open_message_id"`
	OpenID        string `json:"open_id"`
	OpenChatId    string `json:"open_chat_id"`
	Text          string `json:"text_without_at_bot"`
}

type ToolKit struct {
	cli *client.LarkClient
}

func New(c *client.LarkClient) *ToolKit {
	return &ToolKit{c}
}

var pattern *regexp.Regexp = regexp.MustCompile("([^\\s]+)\\s+(.+)\\s*$")

func (t *ToolKit) Do(msg *TextMessageChat) {
	match := pattern.FindSubmatch([]byte(msg.Text))
	if len(match) == 0 {
		return
	}
	var err error
	var rst string

	methods := strings.Split(string(match[1]), ".")
	switch methods[0] {
	case "hash":
		rst, err = doHash(methods, msg.Text)
	default:
		t.cli.MessageSend(msg.OpenChatId, msg.OpenMessageId, "<未知指令>")
		return
	}

	if err != nil {
		t.cli.MessageSend(msg.OpenChatId, msg.OpenMessageId, "ERROR< "+err.Error()+" >")
		return
	}
	t.cli.MessageSend(msg.OpenChatId, msg.OpenMessageId, rst)
}
