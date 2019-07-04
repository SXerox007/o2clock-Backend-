package resetpswd

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"o2clock/collection/forgotpassword"
	"o2clock/constants/collections"
	"o2clock/constants/errormsg"
	"o2clock/db/mongodb"

	"github.com/mongodb/mongo-go-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// change password
type ChangePasswrod struct {
	Email      string
	Phone      string
	Password   string
	Repassword string
}

// get user reset password
func GetUserResetPassword(url string, w http.ResponseWriter, r *http.Request) {
	data, err := GetResetUserPasswordInfo(url)
	if err != nil {
		log.Println("Error in Reset Password:", err)
	}
	OutputHTML(w, "./templates/static/reset_password.html", data)
}

// set user new password
func SetUserNewPassword(w http.ResponseWriter, r *http.Request) {
	log.Println("Request:", r)
	OutputHTML(w, "./templates/static/error_500.html", nil)

}

// output html
func OutputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

/**
*
* Get the reset password password
*
**/
func GetResetUserPasswordInfo(url string) (*forgotpassword.ForgotPasswordAttempt, error) {
	data := &forgotpassword.ForgotPasswordAttempt{}
	filter := bson.M{collections.PARAM_LINK: url}
	mongodb.InitDB()
	res := mongodb.CreateCollection(collections.COLLECTIONS_ALL_LINKS).FindOne(context.Background(), filter)
	//for single data decode
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_DATA_CANT_DECODE, err))
	}
	mongodb.CloseMongoDB()
	return data, nil
}
