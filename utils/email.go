package utils

import (
	"log"
	"net/smtp"
	"os"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/templates"
)

func SendEmail(email string, subject string, lang string, code string, name string, mainText string, buttonText string) {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))

	to := []string{email}

	contentType := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte(
		"To:" + os.Getenv("SMTP_FROM") + "\r\n" +
			"Subject:" + subject + "\r\n" + contentType +
			"\r\n" +
			templates.VerificationEmailTemplate(lang, code, name, mainText, buttonText) + "\r\n")

	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, "api", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
