package forgotpassword

import (
	"io/ioutil"
	"o2clock/collection/allusers"
	"o2clock/core/sendmail"
	"os"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
)

const (
	VERSION = "v1.0"
)

type ForgotPasswordAttempt struct {
	ID        objectid.ObjectID `bson:"_id,omitempty"`
	Email     string            `bson:"email"`
	Version   string            `bson:"version"`
	TimeStamp time.Time         `bson:"timestamp"`
}

func UserForgotPassword(email string) error {
	if err := allusers.ValidateUserEmail(email); err != nil {
		return err
	}
	sender := sendmail.NewSender(os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_PASSWORD"))
	Receiver := []string{email}
	Subject := "Forgot your 02clock password"
	message, err := ioutil.ReadFile("./templates/forgot_password.html")
	if err != nil {
		return err
	}
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, string(message))

	sender.SendMail(Receiver, Subject, bodyMessage)

	return nil

}
