package tasks

import (
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/marveldo/gogin/internal/application/html_temp"
	payload "github.com/marveldo/gogin/internal/application/payloads"
	"github.com/marveldo/gogin/internal/config"
)

func SendEmail(e *payload.EmailPayload) error {
	html_body := html_temp.GetEmailHtml(e)
	addr := "smtp.gmail.com:587"
	from := config.LoadConfig().Email
	to := e.Email
	subject := "Welcome!"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg := "To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		mime + "\r\n" +
		html_body.String()

	auth := sasl.NewPlainClient("", from, config.LoadConfig().EmailPassword)

	err := smtp.SendMail(addr, auth, from, []string{to}, strings.NewReader(msg))
	if err != nil {
		return err
	}
	return nil

}

func SendSuccessMail(e *payload.EmailPayload) error {
	html_body := html_temp.GetPaymentSuccessfulEmail(e)
	addr := "smtp.gmail.com:587"
	from := config.LoadConfig().Email
	to := e.Email 
	subject := "Transaction !"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg := "To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		mime + "\r\n" +
		html_body.String()
	auth := sasl.NewPlainClient("", from, config.LoadConfig().EmailPassword)
	err := smtp.SendMail(addr, auth, from, []string{to}, strings.NewReader(msg))
	if err != nil {
		return err
	}
	return nil
}

func SendFailedEmail(e *payload.EmailPayload)error {
	html_body := html_temp.GetPaymentFailedEmail(e)
	addr := "smtp.gmail.com:587"
	from := config.LoadConfig().Email
	to := e.Email 
	subject := "Transaction !"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg := "To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		mime + "\r\n" +
		html_body.String()
	auth := sasl.NewPlainClient("", from, config.LoadConfig().EmailPassword)
	err := smtp.SendMail(addr, auth, from, []string{to}, strings.NewReader(msg))
	if err != nil {
		return err
	}
	return nil
}