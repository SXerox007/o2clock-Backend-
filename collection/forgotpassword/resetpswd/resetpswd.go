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
	"o2clock/utils"
	"o2clock/utils/pswdmanager"

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

// error
type Error struct {
	Status  int
	Message string
	Title   string
	Style   string
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
	mongodb.InitDB()
	defer mongodb.CloseMongoDB()
	var resp Error
	pass := utils.ParseRequest(r, "password")
	repass := utils.ParseRequest(r, "repassword")
	email := utils.ParseRequest(r, "email")
	resp.Status = http.StatusOK
	resp.Title = "Password Success  :)"
	resp.Message = "Your change password is updated with success. Please login into app with the new password."
	resp.Style = "color:green"
	if pass == repass {
		// set new passsword
		if err := SetNewPassword(email, pswdmanager.GetPswdHash([]byte(pass))); err != nil {
			resp.Status = http.StatusInternalServerError
			resp.Title = errormsg.ERR_MSG_INTERNAL_SERVER
			resp.Message = errormsg.ERR_SRV_NOT_RESPOND
			resp.Style = "color:red"
		}
		OutputHTML(w, "./templates/static/error_500.html", resp)
	} else {
		resp.Status = http.StatusPreconditionRequired
		resp.Title = errormsg.ERR_MSG_PSWD_DECODE
		resp.Message = errormsg.ERR_PASSWORD_NOT_MATCH
		resp.Style = "color:red"
		OutputHTML(w, "./templates/static/error_500.html", resp)
	}
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

// set new password
func SetNewPassword(email, passwordHash string) error {
	filter := bson.M{collections.PARAM_EMAIL: email}
	updateFilter := bson.M{"$set": bson.M{collections.PARAM_PSWD: passwordHash}}
	_, err := mongodb.CreateCollection(collections.COLLECTIONS_ALL_USERS).UpdateOne(context.Background(), filter, updateFilter)
	return err
}
