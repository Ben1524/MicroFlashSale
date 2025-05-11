package utils

import "gopkg.in/gomail.v2"

const fromEmail = "3327869834@qq.com"
const password = "anqdqocycngychca"
const smtpServer = "smtp.qq.com"
const smtpPort = "465"

func SendEmail(toEmail string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpServer, 465, fromEmail, password)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
