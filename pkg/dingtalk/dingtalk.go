package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/levigross/grequests"
)

// BtnOrientation 按钮排列类型
type BtnOrientation int

const (
	// BtnOrientationHorizonal 横排
	BtnOrientationHorizonal = 0
	// BtnOrientationVertical 横排
	BtnOrientationVertical = 1
)

// ActionCardButton ActionCard的按钮内容
type ActionCardButton struct {
	Title     string
	ActionURL string
}

// FeedCard FeedCard的内容
type FeedCard struct {
	Title      string //单条信息文本
	MessageURL string //点击单条信息到跳转链接
	PicURL     string //单条信息后面图片的URL
}

// Client 钉钉消息推送机器人客户端
type Client struct {
	token  string
	secret string
}

// NewClient create a client
func NewClient(token, secret string) *Client {
	return &Client{
		token:  token,
		secret: secret,
	}
}

// SendTextMsg 发送 link类型 消息
// text: 消息内容，如果太长只会部分展示
// atAll: 是否@所有人
// mobiles: 被@人的手机号
func (c *Client) SendTextMsg(text string, atAll bool, mobiles ...string) error {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": text,
		},
		"at": map[string]interface{}{
			"atMobiles": _defaultEmptySlice(mobiles),
			"isAtAll":   atAll,
		},
	}
	// data, err := json.Marshal(msg)
	// log.Printf("%v\n", string(data))
	// return err
	return c.send(msg)
}

// SendLinkMsg 发送 link类型 消息
// title: 消息标题
// text: 消息内容，如果太长只会部分展示
// messageURL: 点击消息跳转的URL
// picURL: 图片URL
// atAll: 是否@所有人
// mobiles: 被@人的手机号
func (c *Client) SendLinkMsg(title, text, messageURL, picURL string, atAll bool, mobiles ...string) error {
	msg := map[string]interface{}{
		"msgtype": "link",
		"link": map[string]interface{}{
			"text":       text,
			"title":      title,
			"messageUrl": messageURL,
			"picUrl":     picURL, //threeExpr(len(picUrl) > 0, picUrl, ""),
		},
		"at": map[string]interface{}{
			"atMobiles": _defaultEmptySlice(mobiles),
			"isAtAll":   atAll,
		},
	}
	return c.send(msg)

}

// SendMarkdownMsg 发送 markdown类型 消息
// title: 首屏会话透出的展示内容
// text: markdown格式的消息
// atAll: 是否@所有人
// mobiles: 被@人的手机号
func (c *Client) SendMarkdownMsg(title, text string, atAll bool, mobiles ...string) error {
	msg := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"text":  text,
			"title": title,
		},
		"at": map[string]interface{}{
			"atMobiles": _defaultEmptySlice(mobiles),
			"isAtAll":   atAll,
		},
	}
	return c.send(msg)
}

// SendSingleActionCardMsg 发送 整体跳转ActionCard类型 消息
// title: 首屏会话透出的展示内容
// text: markdown格式的消息
// singleTitle: 单个按钮的标题
// singleURL: 点击singleTitle按钮触发的URL
// btnOrientation: BtnOrientationHorizonal：按钮竖直排列
//                 BtnOrientationVertical：按钮横向排列
func (c *Client) SendSingleActionCardMsg(title, text, singleTitle, singleURL string, btnOrientation BtnOrientation) error {
	msg := map[string]interface{}{
		"msgtype": "actionCard",
		"actionCard": map[string]interface{}{
			"text":           text,
			"title":          title,
			"singleTitle":    singleTitle,
			"singleURL":      singleURL,
			"btnOrientation": btnOrientation,
		},
	}
	return c.send(msg)
}

// SendActionCardMsg 发送 独立跳转ActionCard类型 消息
// title: 首屏会话透出的展示内容
// text: markdown格式的消息
// hideAvatar: 是否隐藏头像
// btnOrientation: BtnOrientationHorizonal：按钮竖直排列
//                 BtnOrientationVertical：按钮横向排列
// buttons: 按钮
func (c *Client) SendActionCardMsg(title, text string, hideAvatar bool, btnOrientation BtnOrientation, buttons ...ActionCardButton) error {
	actionCard := map[string]interface{}{
		"text":           text,
		"title":          title,
		"hideAvatar":     threeExpr(hideAvatar, "0", "1"),
		"btnOrientation": btnOrientation,
	}

	if buttons != nil && len(buttons) > 0 {
		btns := []map[string]string{}
		for _, b := range buttons {
			btns = append(btns, map[string]string{
				"title":     b.Title,
				"actionURL": b.ActionURL,
			})
		}
		actionCard["btns"] = btns
	}

	msg := map[string]interface{}{
		"msgtype":    "actionCard",
		"actionCard": actionCard,
	}
	return c.send(msg)
}

// SendFeedCardMsg 发送 FeedCard类型 消息
// title: 首屏会话透出的展示内容
// text: markdown格式的消息
// hideAvatar: 是否隐藏头像
// btnOrientation: BtnOrientationHorizonal：按钮竖直排列
//                 BtnOrientationVertical：按钮横向排列
// buttons: 按钮
func (c *Client) SendFeedCardMsg(feedCards []FeedCard) error {
	links := []map[string]string{}

	for _, c := range feedCards {
		links = append(links, map[string]string{
			"title":      c.Title,
			"messageURL": c.MessageURL,
			"picURL":     c.PicURL,
		})
	}

	msg := map[string]interface{}{
		"msgtype": "feedCard",
		"feedCard": map[string]interface{}{
			"links": links,
		},
	}
	return c.send(msg)
}

func (c *Client) send(msg interface{}) error {
	ts := timestamp()
	signStr := sign(ts, c.secret)
	// log.Printf("timestamp=%s, sign=%s", ts, signStr)
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%s&sign=%s",
		c.token, ts, signStr)
	ro := grequests.RequestOptions{
		JSON: msg,
	}
	_, err := grequests.Post(url, &ro)
	return err
	// if err != nil {
	// log.Fatal(err)
	// return err
	// }
	// log.Println(resp.String())
	// return nil
}

func _map(list []interface{}, mapFunc func(interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(list))
	for _, item := range list {
		result = append(result, mapFunc(item))
	}
	return result
}

func threeExpr(cond bool, trueVal interface{}, falseVal interface{}) interface{} {
	if cond {
		return trueVal
	}
	return falseVal
}

func _defaultEmptySlice(mobiles []string) []string {
	if mobiles == nil || len(mobiles) == 0 {
		return []string{}
	}
	return mobiles
}

func timestamp() string {
	return strconv.Itoa(int(time.Now().UnixNano() / 1000000))
}

func sign(timestamp, secret string) string {
	strToSign := timestamp + "\n" + secret
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(strToSign))
	encoded := h.Sum(nil)
	encodedStr := base64.StdEncoding.EncodeToString(encoded)
	escapedEncodedStr := url.PathEscape(encodedStr)
	return escapedEncodedStr
}
