package email

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func newDialer() *gomail.Dialer {
	host := viper.GetString("email.smtp.host")
	port := viper.GetInt("email.smtp.port")
	user := viper.GetString("email.smtp.user")
	password := viper.GetString("email.smtp.password")
	log.Println(host, port, user, password)
	return gomail.NewDialer(host, port, user, password)
}

func newMessage(to string, subject string, body string) *gomail.Message {
	from := viper.GetString("email.smtp.user")
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return m
}

func Send() {
	m := newMessage("quinnn.gao@gmail.com", "Hello!", "Hello <b>Bob</b> and <i>这是一个测试邮件</i>!")
	d := newDialer()
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}

func SendValidationCode(email, code string) error {
	m := newMessage(email, "验证码", fmt.Sprintf("你的验证码是:%s", code))
	return newDialer().DialAndSend(m)
}
