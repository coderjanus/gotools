package dingtalk

import "testing"

var (
	client = NewClient("<YOUR_TOKEN_HERE>",
		"<YOUR_SECRET_HERE>")
	textURL = "https://ding-doc.dingtalk.com/document#/org-dev-guide/custom-robot"
	picURL  = "https://ding-doc.dingtalk.com/document#/org-dev-guide/custom-robot"
)

func TestSendTextMsg(t *testing.T) {

	err := client.SendTextMsg("Hello World!", false)
	if err != nil {
		t.Error(err)
	}
}
func TestSendLinkMsg(t *testing.T) {

	err := client.SendLinkMsg(
		"时代的火车向前开",
		"这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林",
		textURL, "", false)
	if err != nil {
		t.Error(err)
	}
}

func TestSendMarkdownMsg(t *testing.T) {

	err := client.SendMarkdownMsg(
		"杭州天气",
		"#### 杭州天气 @150XXXXXXXX \n> 9度，西北风1级，空气良89，相对温度73%\n> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n> ###### 10点20分发布 [天气](https://www.dingtalk.com) \n",
		false)
	if err != nil {
		t.Error(err)
	}
}

func TestSendSingleActionCardMsg(t *testing.T) {

	err := client.SendSingleActionCardMsg(
		"乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
		`![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png) 
		### 乔布斯 20 年前想打造的苹果咖啡厅 
		Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`,
		"阅读全文",
		textURL,
		BtnOrientationHorizonal)
	if err != nil {
		t.Error(err)
	}
}

func TestSendActionCardMsgH(t *testing.T) {

	err := client.SendActionCardMsg(
		"乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
		`![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png) 
		### 乔布斯 20 年前想打造的苹果咖啡厅 
		Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`,
		false,
		BtnOrientationHorizonal,
		ActionCardButton{
			Title:     "内容不错",
			ActionURL: textURL,
		},
		ActionCardButton{
			Title:     "不感兴趣",
			ActionURL: textURL,
		},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestSendActionCardMsgV(t *testing.T) {

	err := client.SendActionCardMsg(
		"乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
		`![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png) 
		### 乔布斯 20 年前想打造的苹果咖啡厅 
		Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`,
		false,
		BtnOrientationVertical,
		ActionCardButton{
			Title:     "内容不错",
			ActionURL: textURL,
		},
		ActionCardButton{
			Title:     "不感兴趣",
			ActionURL: textURL,
		},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestSendFeedCardMsg(t *testing.T) {
	err := client.SendFeedCardMsg(
		[]FeedCard{
			{
				Title:      "时代的火车向前开1",
				MessageURL: textURL,
				PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
			}, {
				Title:      "时代的火车向前开2",
				MessageURL: textURL,
				PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
			}, {
				Title:      "时代的火车向前开3",
				MessageURL: textURL,
				PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
			},
		},
	)
	if err != nil {
		t.Error(err)
	}
}
