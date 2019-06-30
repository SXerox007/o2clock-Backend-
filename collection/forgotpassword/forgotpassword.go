package forgotpassword

import (
	"context"
	"fmt"
	"io/ioutil"
	"o2clock/collection/allusers"
	"o2clock/constants/appconstant"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/core/sendmail"
	"o2clock/db/mongodb"
	"o2clock/utils"
	"os"
	"strings"
	"time"

	objectid "github.com/mongodb/mongo-go-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	VERSION = "v1.0"
)

type ForgotPasswordAttempt struct {
	ID        objectid.ObjectID `bson:"_id,omitempty"`
	Email     string            `bson:"email"`
	Phone     string            `bson:"phone"`
	ExtId     objectid.ObjectID `bson:"ext_id"`
	Link      string            `bson:"link"`
	Event     int               `bson:"event"`
	State     int               `bson:"state"`
	Version   string            `bson:"version"`
	TimeStamp time.Time         `bson:"timestamp"`
}

func UserForgotPassword(email string) error {
	data, err := allusers.ValidateUserEmail(email)
	if err != nil {
		return err
	}
	sender := sendmail.NewSender(os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_PASSWORD"))
	Receiver := []string{email}
	Subject := "Forgot your 02clock password"
	message, err := ioutil.ReadFile("./templates/forgot_password.html")
	if err != nil {
		return err
	}
	var link = "http://localhost:5051/v1/user/setpassword/" + utils.RandomStringGenerateWithType(32, appconstant.ALPHA_NUM)
	var replacer = strings.NewReplacer("text-link-forgot", "<a href='"+link+"' target='_blank'>"+link+"</a>", "button-link-forgot", "href='"+link+"'")
	output := replacer.Replace(string(message))
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, output)
	if err = sender.SendMail(Receiver, Subject, bodyMessage); err != nil {
		return err
	}
	info := ForgotPasswordAttempt{
		Email:     data.Email,
		Phone:     data.Phone,
		ExtId:     data.ID,
		Link:      link,
		Event:     appconstant.EVENT_FORGOT_PSWD,
		State:     appconstant.STATE_LINK_GENERATE,
		Version:   VERSION,
		TimeStamp: time.Now(),
	}
	//save link with timestamp and user info
	res, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_LINKS).InsertOne(context.Background(), info)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL, err))
	}

	_, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_INTERNAL_OID, ok))
	}
	return nil

}
