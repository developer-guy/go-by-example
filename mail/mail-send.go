package mail

import (
	"net/smtp"
	"bytes"
	"html/template"
	"log"
	"strconv"
	"os"
)

type EmailUser struct {
	Username    string
	Password    string
	Emailserver string
	Port        int
}

type MailTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,

{{.From}}
`

func SendMail(done chan bool) {
	emailUser := &EmailUser{"bapaydin67", os.Getenv("PWD"), "smtp.gmail.com", 587}
	auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.Emailserver)
	var err error
	var doc bytes.Buffer

	mailTemplateData := &MailTemplateData{
		"SmtpEmailSender",
		"bapaydin67@gmail.com",
		"This is the e-mail subject line!",
		"Hello, this is a test e-mail body.",
	}

	t := template.New("emailTemplate")
	t, err = t.Parse(emailTemplate)

	if err != nil {
		log.Print("error trying to parse mail template")
		done <- false
	}

	err = t.Execute(&doc, mailTemplateData)

	if err != nil {
		log.Print("error trying to execute mail template")
		done <- false
	}

	err = smtp.SendMail(emailUser.Emailserver+":"+strconv.Itoa(emailUser.Port),
		auth,
		emailUser.Username,
		[]string{"bapaydin67@gmail.com"},
		doc.Bytes())

	if err != nil {
		log.Print("ERROR: attempting to send a mail ", err)
		done <- false
	}

	done <- true
}
