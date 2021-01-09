package mail

import (
	"errors"
	"time"

	gomail "gopkg.in/gomail.v2"
)

// Conf mail config
type Conf struct {
	Server   string
	Port     int
	User     string
	Password string
	From     string
	IdelTime time.Duration
}

// Client email client
type Client struct {
	conf *Conf
}

// NewClient create a new email client
func NewClient(conf Conf) *Client {
	return &Client{conf: &conf}
}

// SendTextMail send text email
func (c *Client) SendTextMail(subject, content string, to []string, cc ...string) error {
	return c.send("text/plain", subject, content, to, nil, cc...)
}

// SendTextMailWithAttachments send text email
func (c *Client) SendTextMailWithAttachments(subject, content string, to []string, attachments []string, cc ...string) error {
	return c.send("text/plain", subject, content, to, attachments, cc...)
}

// SendHTMLMail send HTML email
func (c *Client) SendHTMLMail(subject, content string, to []string, cc ...string) error {
	return c.send("text/html", subject, content, to, nil, cc...)
}

// SendHTMLMailWithAttachments send HTML email with attachments
func (c *Client) SendHTMLMailWithAttachments(subject, content string, to []string, attachments []string, cc ...string) error {
	return c.send("text/html", subject, content, to, attachments, cc...)
}

func (c *Client) send(contentType, subject, content string, to []string, attachments []string, cc ...string) error {
	if len(subject) == 0 {
		return errors.New("subject cannot be empty")
	}
	if len(content) == 0 {
		return errors.New("content cannot be empty")
	}
	if to == nil || len(to) == 0 {
		return errors.New("to cannot be empty")
	}
	d := gomail.NewDialer(c.conf.Server, c.conf.Port, c.conf.User, c.conf.Password)
	m := gomail.NewMessage()

	m.SetHeader("From", c.conf.From)
	m.SetHeader("To", to...)
	if cc != nil && len(cc) > 0 {
		m.SetHeader("Cc", cc...)
	}
	m.SetHeader("Subject", subject)
	m.SetBody(contentType, content)
	if nil != attachments && len(attachments) > 0 {
		for _, a := range attachments {
			m.Attach(a)
		}
	}

	err := d.DialAndSend(m)
	return err
}
