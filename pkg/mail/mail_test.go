package mail

import "testing"

func newClient() *Client {
	conf := Conf{
		Server:   "smtp.exmail.qq.com",
		Port:     25,
		User:     "dev@estockapp.com",
		Password: "Fnn4X4dfsSBvcA9t",
		From:     "dev@estockapp.com",
	}
	return NewClient(conf)
}
func TestSendTextMsg(t *testing.T) {
	client := newClient()
	err := client.SendTextMail("Test", "Test from gomail client", []string{"janushuang@163.com", "janus.huang.cn@gmail.com"}, "monkyhuang@163.com")
	if err != nil {
		t.Error(err)
	}
}
func TestSendTextWithAttachmentMsg(t *testing.T) {
	client := newClient()
	err := client.SendTextMailWithAttachments("Test", "Test from gomail client", []string{"janushuang@163.com"}, []string{"./mail_test.go"})
	if err != nil {
		t.Error(err)
	}
}
func TestSendHTMLMsg(t *testing.T) {
	client := newClient()
	content := `
	<html>
		<body><H1>Hello World!</H1></body
	</html>
	`
	err := client.SendHTMLMail("Test Html", content, []string{"janushuang@163.com"})
	if err != nil {
		t.Error(err)
	}
}

func TestSendHTMLWithAttachmentMsg(t *testing.T) {
	client := newClient()
	content := `
	<html>
		<body><H1>Hello World!</H1></body
	</html>
	`
	err := client.SendHTMLMailWithAttachments("Test Html", content, []string{"janushuang@163.com"}, []string{"./mail_test.go"})
	if err != nil {
		t.Error(err)
	}
}
