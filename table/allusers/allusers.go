package allusers

import (
	"fmt"
	"o2clock/api-proto/onboarding/register"
	"o2clock/constants/errormsg"
	db "o2clock/db/postgres"
	"o2clock/table/accesstoken"
	"o2clock/utils/pswdmanager"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	VERSION                   = "v1.0"
	SQL_STATEMENT_CREATE_USER = `
	INSERT INTO all_users (phone, first_name, last_name,company_name,
	country_code,user_name,password,email,lat,lan,version,capture_time)
	VALUES ($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	RETURNING id`
	SQL_STATEMENT_GET_USER_DATA_USING_USER_NAME = `SELECT * FROM all_users WHERE user_name=$1;`
	SQL_STATEMENT_GET_USER_DATA_USING_EMAIL     = `SELECT * FROM all_users WHERE email=$1;`
)

type Users struct {
	TableName   struct{}  `sql:"all_users" json:"-"`
	Id          string    `param:"id"`
	Phone       string    `param:"phone"`
	FirstName   string    `param:"first_name"`
	LastName    string    `param:"last_name"`
	CompanyName string    `param:"company_name"`
	CountryCode string    `param:"country_code"`
	UserName    string    `param:"user_name"`
	Password    string    `param:"password"`
	Email       string    `param:"email"`
	Lat         float64   `param:"lat"`
	Lan         float64   `param:"lan"`
	Version     string    `param:"version"`
	CaptureTime time.Time `param:"capture_time"`
}

func CreateUser(req *regsiterpb.RegisterUserRequest) (string, error) {
	if err := validateUser(req); err != nil {
		return "", nil
	}
	sqlStatement := SQL_STATEMENT_CREATE_USER
	var id string
	err := db.GetClient().QueryRow(sqlStatement, req.GetPhone(), req.GetFirstName(),
		req.GetLastName(), req.GetCompanyName(), req.GetCountryCode(), req.GetUserName(),
		pswdmanager.GetPswdHash([]byte(req.GetPassword())), req.GetEmail(), req.GetLocation().GetLat(),
		req.GetLocation().GetLan(), VERSION, time.Now()).Scan(&id)
	if err != nil {
		return "", status.Errorf(
			codes.Internal,
			fmt.Sprintln(errormsg.ERR_MSG_INTERNAL_SERVER, err))
	}
	return accesstoken.CreateUserAccessToken(id, req.GetUserName())

}

func validateUser(req *regsiterpb.RegisterUserRequest) error {
	sqlStatement := SQL_STATEMENT_GET_USER_DATA_USING_USER_NAME
	//var user Users
	row := db.GetClient().QueryRow(sqlStatement, req.GetUserName())
	if row != nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_USERNAME_ALREADY_EXIST))
	}
	sqlStatement = SQL_STATEMENT_GET_USER_DATA_USING_EMAIL
	//var user Users
	row = db.GetClient().QueryRow(sqlStatement, req.GetEmail())
	if row != nil {
		return status.Errorf(
			codes.Aborted,
			fmt.Sprintln(errormsg.ERR_MSG_EMAIL_ALREADY_EXIST))
	}
	return nil
}
