package email

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var (
	EMAIL_SMTP_HOST = viper.GetString("EMAIL_SMTP_HOST")
	EMAIL_SMTP_PORT = viper.GetString("EMAIL_SMTP_PORT")
	EMAIL_USER      = viper.GetString("EMAIL_USER")
	EMAIL_PWD       = viper.GetString("EMAIL_PWD")
)

func Send() {
	host := viper.GetString("email.smtp.host")
	port := viper.GetInt("email.smtp.port")
	user := viper.GetString("email.smtp.user")
	password := viper.GetString("email.smtp.password")
	m := gomail.NewMessage()

	m.SetHeader("From", "qiangqinag_gao@foxmail.com")
	m.SetHeader("To", "quinnn.gao@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer(host, port, user, password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
