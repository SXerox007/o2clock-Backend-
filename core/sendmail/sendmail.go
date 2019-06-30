package sendmail

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"o2clock/constants/appconstant"
	"strings"
)

const (
	SMTPServer = "smtp.gmail.com"
)

type Sender struct {
	User     string
	Password string
}

func NewSender(Username, Password string) Sender {

	return Sender{Username, Password}
}

func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) error {

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return err
	}

	fmt.Println("Mail sent successfully!")
	return nil
}

func (sender Sender) WriteEmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header[appconstant.EMAIL_FROM] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header[appconstant.EMAIL_TO] = receipient
	header[appconstant.EMAIL_SUBJECT] = subject
	header[appconstant.MIME_VERSION] = appconstant.VERSION
	header[appconstant.CONTENT_TYPE] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header[appconstant.CONTENT_TRANSFER_ENCODING] = appconstant.QUOTED_PRINT
	header[appconstant.CONTENT_DISPOSITION] = appconstant.INLINE

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *Sender) WriteHTMLEmail(dest []string, subject, bodyMessage string) string {
	return sender.WriteEmail(dest, appconstant.TXT_HTML, subject, bodyMessage)
}

func (sender *Sender) WritePlainEmail(dest []string, subject, bodyMessage string) string {
	return sender.WriteEmail(dest, appconstant.TXT_PLAIN, subject, bodyMessage)
}
